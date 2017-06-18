// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: github.com/overcyn/matcha/pb/view/tabnav/tabnav.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers.h>
#else
 #import "GPBProtocolBuffers.h"
#endif

#if GOOGLE_PROTOBUF_OBJC_VERSION < 30002
#error This file was generated by a newer version of protoc which is incompatible with your Protocol Buffer library sources.
#endif
#if 30002 < GOOGLE_PROTOBUF_OBJC_MIN_SUPPORTED_VERSION
#error This file was generated by an older version of protoc which is incompatible with your Protocol Buffer library sources.
#endif

// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

CF_EXTERN_C_BEGIN

@class MatchaPBImage;
@class MatchaTabScreenPBChildView;

NS_ASSUME_NONNULL_BEGIN

#pragma mark - MatchaTabScreenPBViewRoot

/**
 * Exposes the extension registry for this file.
 *
 * The base class provides:
 * @code
 *   + (GPBExtensionRegistry *)extensionRegistry;
 * @endcode
 * which is a @c GPBExtensionRegistry that includes all the extensions defined by
 * this file and all files that it depends on.
 **/
@interface MatchaTabScreenPBViewRoot : GPBRootObject
@end

#pragma mark - MatchaTabScreenPBChildView

typedef GPB_ENUM(MatchaTabScreenPBChildView_FieldNumber) {
  MatchaTabScreenPBChildView_FieldNumber_Id_p = 1,
  MatchaTabScreenPBChildView_FieldNumber_Title = 2,
  MatchaTabScreenPBChildView_FieldNumber_Icon = 3,
  MatchaTabScreenPBChildView_FieldNumber_SelectedIcon = 4,
  MatchaTabScreenPBChildView_FieldNumber_Badge = 5,
};

@interface MatchaTabScreenPBChildView : GPBMessage

@property(nonatomic, readwrite) int64_t id_p;

@property(nonatomic, readwrite, copy, null_resettable) NSString *title;

@property(nonatomic, readwrite, strong, null_resettable) MatchaPBImage *icon;
/** Test to see if @c icon has been set. */
@property(nonatomic, readwrite) BOOL hasIcon;

@property(nonatomic, readwrite, strong, null_resettable) MatchaPBImage *selectedIcon;
/** Test to see if @c selectedIcon has been set. */
@property(nonatomic, readwrite) BOOL hasSelectedIcon;

@property(nonatomic, readwrite, copy, null_resettable) NSString *badge;

@end

#pragma mark - MatchaTabScreenPBView

typedef GPB_ENUM(MatchaTabScreenPBView_FieldNumber) {
  MatchaTabScreenPBView_FieldNumber_ScreensArray = 1,
  MatchaTabScreenPBView_FieldNumber_SelectedIndex = 2,
  MatchaTabScreenPBView_FieldNumber_EventFunc = 3,
};

@interface MatchaTabScreenPBView : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<MatchaTabScreenPBChildView*> *screensArray;
/** The number of items in @c screensArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger screensArray_Count;

@property(nonatomic, readwrite) int64_t selectedIndex;

@property(nonatomic, readwrite) int64_t eventFunc;

@end

#pragma mark - MatchaTabScreenPBEvent

typedef GPB_ENUM(MatchaTabScreenPBEvent_FieldNumber) {
  MatchaTabScreenPBEvent_FieldNumber_SelectedIndex = 1,
};

@interface MatchaTabScreenPBEvent : GPBMessage

@property(nonatomic, readwrite) int64_t selectedIndex;

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
