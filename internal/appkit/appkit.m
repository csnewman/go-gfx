#import <Foundation/Foundation.h>
#import <AppKit/AppKit.h>
#import <QuartzCore/CAMetalLayer.h>
#include "appkit.h"

@interface GfxApplicationDelegate : NSObject <NSApplicationDelegate>
@end

@implementation GfxApplicationDelegate
- (void)applicationDidFinishLaunching:(NSNotification *)notification {
    [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];

    gfx_ak_init_callback();
}

- (void)stubThread:(id)sender {
}
@end

int gfx_ak_run() {
    @autoreleasepool {
        if (![NSThread isMainThread]) {
            return GFX_NOT_MAIN_THREAD;
        }

        [NSApplication sharedApplication];

        GfxApplicationDelegate *appDelegate = [[[GfxApplicationDelegate alloc] init] autorelease];

        // Ensure we are in multi-threading mode
        [NSThread detachNewThreadSelector:@selector(stubThread:)
                                 toTarget:appDelegate
                               withObject:nil];

        [NSApp setDelegate:appDelegate];

        [NSApp run];

        return GFX_SUCCESS;
    }
}

void gfx_ak_stop() {
    [NSApp stop:NSApp];
}

@class GfxWindow;
@class GfxWindowDelegate;
@class GfxView;

@interface GfxWindowContext : NSObject {
@public
    uint64_t wid;
    GfxWindow *window;
    GfxWindowDelegate *delegate;
    GfxView *view;
    CAMetalLayer *layer;
}

- (instancetype)initWithWID:(uint64_t)wid;
@end

@implementation GfxWindowContext
- (instancetype)initWithWID:(uint64_t)pwid {
    self = [super init];
    if (self != nil)
        self->wid = pwid;

    return self;
}
@end

@interface GfxWindow : NSWindow
@end

@implementation GfxWindow
- (BOOL)canBecomeKeyWindow {
    return YES;
}

- (BOOL)canBecomeMainWindow {
    return YES;
}
@end

@interface GfxWindowDelegate : NSObject <NSWindowDelegate> {
    GfxWindowContext *context;
}

- (instancetype)initWithContext:(GfxWindowContext *)ctx;
@end

@implementation GfxWindowDelegate
- (instancetype)initWithContext:(GfxWindowContext *)ctx {
    self = [super init];
    if (self != nil)
        context = ctx;

    return self;
}

- (BOOL)windowShouldClose:(NSWindow *)sender {
    gfx_ak_close_requested_callback(context->wid);

    return NO;
}

- (void)windowWillClose:(NSNotification *)notification {
    gfx_ak_window_closed_callback(context->wid);
}

@end

@interface GfxView : NSView <NSTextInputClient, CALayerDelegate> {
    GfxWindowContext *context;
}

- (instancetype)initWithContext:(GfxWindowContext *)ctx;

@end

@implementation GfxView
- (instancetype)initWithContext:(GfxWindowContext *)ctx {
    self = [super init];
    if (self != nil)
        context = ctx;

    [self setWantsLayer:YES];
    [self setLayerContentsRedrawPolicy:NSViewLayerContentsRedrawDuringViewResize];

    return self;
}

- (CALayer *)makeBackingLayer {
    context->layer = [CAMetalLayer layer];
    [context->layer setDelegate:self];
    return context->layer;
}

- (BOOL)canBecomeKeyView {
    return YES;
}

- (BOOL)acceptsFirstResponder {
    return YES;
}

- (BOOL)wantsUpdateLayer {
    return YES;
}

 - (void)updateLayer {
     gfx_ak_draw_callback(context->wid);
 }

- (void)displayLayer:(CALayer *)layer {
    gfx_ak_draw_callback(context->wid);
}

- (BOOL)canDrawSubviewsIntoLayer {
    return NO;
}

- (BOOL)hasMarkedText {
    return NO;
}

- (NSRange)markedRange {
    return NSMakeRange(NSNotFound, 0);
}

- (NSRange)selectedRange {
    return NSMakeRange(NSNotFound, 0);
}

- (NSRect)firstRectForCharacterRange:(NSRange)range actualRange:(nullable NSRangePointer)actualRange {
    return NSMakeRect(0, 0, 0, 0);
}

- (NSUInteger)characterIndexForPoint:(NSPoint)point {
    return 0;
}

- (NSArray<NSAttributedStringKey> *)validAttributesForMarkedText {
    return [NSArray array];
}

- (NSAttributedString *)attributedSubstringForProposedRange:(NSRange)range actualRange:(nullable NSRangePointer)actualRange {
    return nil;
}

- (void)insertText:(nonnull id)string replacementRange:(NSRange)replacementRange {
}


- (void)setMarkedText:(nonnull id)string selectedRange:(NSRange)selectedRange replacementRange:(NSRange)replacementRange {
}

- (void)unmarkText {
}

@end

int gfx_ak_new_window(uint64_t wid, int width, int height, id *res) {
    @autoreleasepool {
        GfxWindowContext *ctx = [[GfxWindowContext alloc] initWithWID:wid];
        *res = ctx;

        NSRect contentRect = NSMakeRect(0, 0, width, height);
        NSUInteger styleMask = NSWindowStyleMaskMiniaturizable | NSWindowStyleMaskTitled | NSWindowStyleMaskClosable |
                               NSWindowStyleMaskResizable;

        ctx->window = [[GfxWindow alloc]
                initWithContentRect:contentRect
                          styleMask:styleMask
                            backing:NSBackingStoreBuffered
                              defer:NO
        ];
        [ctx->window setReleasedWhenClosed:NO];

        ctx->delegate = [[GfxWindowDelegate alloc] initWithContext:ctx];
        [ctx->window setDelegate:ctx->delegate];

        ctx->view = [[GfxView alloc] initWithContext:ctx];
        [ctx->window setContentView:ctx->view];
        [ctx->window makeFirstResponder:ctx->view];
        [ctx->view setNeedsDisplay:YES];

        [ctx->window setTitle:@"hello"];
        [ctx->window setRestorable:NO];
        [ctx->window setTabbingMode:NSWindowTabbingModeDisallowed];
        [ctx->window setCollectionBehavior:(NSWindowCollectionBehaviorFullScreenPrimary |
                                            NSWindowCollectionBehaviorManaged)];
        [ctx->window setAcceptsMouseMovedEvents:YES];

        [ctx->window center];
        [ctx->window makeKeyAndOrderFront:NSApp];
        [ctx->window orderFrontRegardless];

        return GFX_SUCCESS;
    }
}

void gfx_ak_close_window(id w) {
    @autoreleasepool {
        GfxWindowContext *ctx = w;

        [ctx->window close];
    }
}

void gfx_ak_free_context(id w) {
    @autoreleasepool {
        GfxWindowContext *ctx = w;

        [ctx->window setDelegate:nil];
        [ctx->delegate release];
        [ctx->window setContentView:nil];
        [ctx->view release];
        [ctx->window release];
        [ctx release];
    }
}
