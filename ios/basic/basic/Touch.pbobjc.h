// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: touch.proto

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

@class GPBAny;
@class MochiPBRecognizer;

NS_ASSUME_NONNULL_BEGIN

#pragma mark - MochiPBTouchRoot

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
@interface MochiPBTouchRoot : GPBRootObject
@end

#pragma mark - MochiPBRecognizer

typedef GPB_ENUM(MochiPBRecognizer_FieldNumber) {
  MochiPBRecognizer_FieldNumber_Id_p = 1,
  MochiPBRecognizer_FieldNumber_Name = 2,
  MochiPBRecognizer_FieldNumber_Recognizer = 3,
};

@interface MochiPBRecognizer : GPBMessage

@property(nonatomic, readwrite) int64_t id_p;

@property(nonatomic, readwrite, copy, null_resettable) NSString *name;

@property(nonatomic, readwrite, strong, null_resettable) GPBAny *recognizer;
/** Test to see if @c recognizer has been set. */
@property(nonatomic, readwrite) BOOL hasRecognizer;

@end

#pragma mark - MochiPBRecognizerList

typedef GPB_ENUM(MochiPBRecognizerList_FieldNumber) {
  MochiPBRecognizerList_FieldNumber_RecognizersArray = 1,
};

@interface MochiPBRecognizerList : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<MochiPBRecognizer*> *recognizersArray;
/** The number of items in @c recognizersArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger recognizersArray_Count;

@end

#pragma mark - MochiPBTapRecognizer

typedef GPB_ENUM(MochiPBTapRecognizer_FieldNumber) {
  MochiPBTapRecognizer_FieldNumber_Count = 1,
  MochiPBTapRecognizer_FieldNumber_RecognizedFunc = 2,
};

@interface MochiPBTapRecognizer : GPBMessage

@property(nonatomic, readwrite) int64_t count;

@property(nonatomic, readwrite) int64_t recognizedFunc;

@end

#pragma mark - MochiPBPressRecognizer

@interface MochiPBPressRecognizer : GPBMessage

@end

#pragma mark - MochiPBPanRecognizer

@interface MochiPBPanRecognizer : GPBMessage

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
