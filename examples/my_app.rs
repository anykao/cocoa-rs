#[warn(unused_imports)]
extern crate cocoa;
extern crate block;
#[macro_use]
extern crate objc;

use block::ConcreteBlock;

use cocoa::base::{selector, nil, class, id, NO};
use cocoa::foundation::{NSUInteger,
                        NSRect,
                        NSPoint,
                        NSSize,
                        NSAutoreleasePool,
                        NSProcessInfo,
                        NSString};
use cocoa::appkit::{NSApp,
                    NSApplication,
                    NSApplicationActivationPolicyAccessory,
                    NSApplicationActivationPolicyRegular,
                    NSWindow,
                    NSTitledWindowMask,
                    NSBackingStoreBuffered,
                    NSMenu,
                    NSMenuItem,
                    NSImage,
                    NSVariableStatusItemLength, NSStatusBar,
                    NSStatusItem,
//                    NSButton,
                    NSEvent,
                    NSAlternateKeyMask,
                    NSRunningApplication,
                    NSApplicationActivateIgnoringOtherApps};

use cocoa::appkit::NSKeyDownMask;
fn main() {
    unsafe {
        let _pool = NSAutoreleasePool::new(nil);
        let app = NSApp();
        app.setActivationPolicy_(NSApplicationActivationPolicyRegular);
        let status_item = NSStatusBar::systemStatusBar(nil).statusItemWithLength_(NSVariableStatusItemLength).autorelease();
        let img = NSString::alloc(nil).init_str("examples/radio@2x.png");
        let icon = NSImage::alloc(nil).initWithContentsOfFile_(img);
        {
            use cocoa::appkit::NSButton;
            status_item.button().setImage_(icon);
        }

        let quit_prefix = NSString::alloc(nil).init_str("Quit ");
        let quit_title = quit_prefix.stringByAppendingString_(
            NSProcessInfo::processInfo(nil).processName()
        );
        let quit_action = selector("terminate:");
        let quit_key = NSString::alloc(nil).init_str("q");
        let quit_item = NSMenuItem::alloc(nil).initWithTitle_action_keyEquivalent_(
            quit_title,
            quit_action,
            quit_key
        ).autorelease();
        let menu = NSMenu::new(nil).autorelease();
        menu.addItem_(quit_item);
        status_item.setMenu_(menu);

        let window = NSWindow::alloc(nil)
            .initWithContentRect_styleMask_backing_defer_(NSRect::new(NSPoint::new(0., 0.),
                                                                      NSSize::new(800., 600.)),
                                                          NSTitledWindowMask,
                                                          NSBackingStoreBuffered,
                                                          NO)
            .autorelease();
        window.cascadeTopLeftFromPoint_(NSPoint::new(20., 20.));
        window.center();
        let title = NSString::alloc(nil).init_str("Hello World!");
        window.setTitle_(title);
        window.makeKeyAndOrderFront_(nil);

        let content = create_table();
        window.setContentView_(content);

        let block = ConcreteBlock::new(move |event: id| {
            if event.keyCode() == 12 && event.modifierFlags().contains(NSAlternateKeyMask) {
                app.stop_(nil);
            }
            nil
        });
        let block = & *block.copy();

        NSEvent::addLocalMonitorForEventsMatchingMask_handler_(nil, NSKeyDownMask, block);

        let current_app = NSRunningApplication::currentApplication(nil);
        current_app.activateWithOptions_(NSApplicationActivateIgnoringOtherApps);
        app.run();
    }
}

fn create_table() -> id {
    unsafe {
        use cocoa::appkit::{NSTableView, NSColor};
        // create a tab View
        let tableView = NSTableView::new(nil)
            .initWithFrame_(NSRect::new(NSPoint::new(0., 0.), NSSize::new(200., 200.)));
        NSTableView::setBackgroundColor_(tableView, NSColor::clearColor(nil));
        tableView
    }
}
