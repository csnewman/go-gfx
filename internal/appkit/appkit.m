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
@end

@interface GfxWindow : NSWindow {}
@end

@implementation GfxWindow
- (BOOL)canBecomeKeyWindow {
    return YES;
}

- (BOOL)canBecomeMainWindow {
    return YES;
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
        [NSApp setDelegate:appDelegate];

        [NSApp run];

        return GFX_SUCCESS;
    }
}

GfxWindow *window;

int gfx_ak_new_window(int width, int height) {
    @autoreleasepool {
        NSRect contentRect = NSMakeRect(0, 0, width, height);
        NSUInteger styleMask = NSWindowStyleMaskMiniaturizable | NSWindowStyleMaskTitled | NSWindowStyleMaskClosable |
                               NSWindowStyleMaskResizable;

        window = [[GfxWindow alloc]
                initWithContentRect:contentRect
                          styleMask:styleMask
                            backing:NSBackingStoreBuffered
                              defer:NO
        ];

        // todo: setcontentview, makefirstresponder setDelegate

        [window setTitle:@"hello"];
        [window setRestorable:NO];
        [window setTabbingMode:NSWindowTabbingModeDisallowed];
        [window setCollectionBehavior:(NSWindowCollectionBehaviorFullScreenPrimary |
                                       NSWindowCollectionBehaviorManaged)];
        [window setAcceptsMouseMovedEvents:YES];

        [window center];
        [window makeKeyAndOrderFront:NSApp];
        [window orderFrontRegardless];

        return GFX_SUCCESS;
    }
}
