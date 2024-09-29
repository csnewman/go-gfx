#import <Foundation/Foundation.h>
#import <AppKit/AppKit.h>
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

GfxApplicationDelegate *appDelegate;

int gfx_ak_run() {
    @autoreleasepool {
        if (![NSThread isMainThread]) {
            return GFX_NOT_MAIN_THREAD;
        }

        [NSApplication sharedApplication];

        appDelegate = [[GfxApplicationDelegate alloc] init];

        // Ensure we are in multi-threading mode
        [NSThread detachNewThreadSelector:@selector(stubThread:)
                                 toTarget:appDelegate
                               withObject:nil];

        [NSApp setDelegate:appDelegate];

        [NSApp run];

        return GFX_SUCCESS;
    }
}

@class GfxWindow;

@interface GfxWindowContext : NSObject {
@public
    uint32_t wid;
    GfxWindow *window;
}

- (instancetype)initWithWID:(uint32_t)wid;
@end

@implementation GfxWindowContext
- (instancetype)initWithWID:(uint32_t)pwid {
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
    [context->window setDelegate:nil];

    gfx_ak_window_closed_callback(context->wid);

    [self release];
}

@end

int gfx_ak_new_window(uint32_t wid, int width, int height, id *res) {
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


        GfxWindowDelegate *delegate = [[GfxWindowDelegate alloc] initWithContext:ctx];
        [ctx->window setDelegate:delegate];

        // todo: setcontentview, makefirstresponder

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
