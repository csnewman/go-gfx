#import <Foundation/Foundation.h>
#import <AppKit/AppKit.h>
#import <QuartzCore/CAMetalLayer.h>
#import <QuartzCore/CAMetalDisplayLink.h>
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

@interface GfxView : NSView <NSTextInputClient, CALayerDelegate, CAMetalDisplayLinkDelegate> {
    GfxWindowContext *context;
    CAMetalDisplayLink *displayLink;
    CFTimeInterval _previousTargetPresentationTimestamp;
}

- (instancetype)initWithContext:(GfxWindowContext *)ctx;

- (void)stopMetalLink;

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
    [context->view stopMetalLink];

    gfx_ak_window_closed_callback(context->wid);
}

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
    context->layer = [[CAMetalLayer layer] retain];
    [context->layer setDelegate:self];
    return context->layer;
}

- (void)viewDidMoveToWindow {
    [displayLink invalidate];

    displayLink = [[CAMetalDisplayLink alloc] initWithMetalLayer:context->layer];
//     displayLink.preferredFrameRateRange = CAFrameRateRangeMake(120.0, 120.0, 120.0);
    displayLink.preferredFrameLatency = 2;
    displayLink.paused = NO;
    displayLink.delegate = self;

    _previousTargetPresentationTimestamp = CACurrentMediaTime();
    [displayLink addToRunLoop:[NSRunLoop currentRunLoop] forMode:NSDefaultRunLoopMode];
    [displayLink addToRunLoop:[NSRunLoop currentRunLoop] forMode:NSEventTrackingRunLoopMode];

    [self resizeDrawable];
}

- (void)metalDisplayLink:(CAMetalDisplayLink *)link
             needsUpdate:(CAMetalDisplayLinkUpdate *_Nonnull)update {
    CFTimeInterval deltaTime = _previousTargetPresentationTimestamp - update.targetPresentationTimestamp;
    _previousTargetPresentationTimestamp = update.targetPresentationTimestamp;

    gfx_ak_draw_callback(context->wid, update.drawable);
}

- (void)stopMetalLink {
    [displayLink removeFromRunLoop:[NSRunLoop mainRunLoop] forMode:NSDefaultRunLoopMode];
    [displayLink removeFromRunLoop:[NSRunLoop mainRunLoop] forMode:NSEventTrackingRunLoopMode];
}

- (void)dealloc {
    [displayLink invalidate];

    [super dealloc];
}

- (void)viewDidChangeBackingProperties {
    [super viewDidChangeBackingProperties];
    [self resizeDrawable];
}

- (void)setFrameSize:(NSSize)size {
    [super setFrameSize:size];
    [self resizeDrawable];
}

- (void)setBoundsSize:(NSSize)size {
    [super setBoundsSize:size];
    [self resizeDrawable];
}

- (void)resizeDrawable {
    CGFloat scaleFactor = self.window.screen.backingScaleFactor;
    CGSize newSize = self.bounds.size;
    newSize.width *= scaleFactor;
    newSize.height *= scaleFactor;

    if (newSize.width <= 0 || newSize.height <= 0) {
        return;
    }

    if (newSize.width == context->layer.drawableSize.width &&
        newSize.height == context->layer.drawableSize.height) {
        return;
    }

    context->layer.drawableSize = newSize;

    gfx_ak_resize_callback(context->wid, newSize.width, newSize.height);
}

- (BOOL)canBecomeKeyView {
    return YES;
}

- (BOOL)acceptsFirstResponder {
    return YES;
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

int gfx_ak_new_window(uint64_t wid, int width, int height, id *res, id *res_wh) {
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

        *res_wh = ctx->layer;

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
