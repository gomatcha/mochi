// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gomatcha.io/matcha/pb/view/alert/alert.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers_RuntimeSupport.h>
#else
 #import "GPBProtocolBuffers_RuntimeSupport.h"
#endif

 #import "gomatcha.io/matcha/pb/view/alert/Alert.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

#pragma mark - MatchaAlertPBAlertRoot

@implementation MatchaAlertPBAlertRoot

// No extensions in the file and no imports, so no need to generate
// +extensionRegistry.

@end

#pragma mark - MatchaAlertPBAlertRoot_FileDescriptor

static GPBFileDescriptor *MatchaAlertPBAlertRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@"matcha.view.alert"
                                                 objcPrefix:@"MatchaAlertPB"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - Enum MatchaAlertPBButtonStyle

GPBEnumDescriptor *MatchaAlertPBButtonStyle_EnumDescriptor(void) {
  static GPBEnumDescriptor *descriptor = NULL;
  if (!descriptor) {
    static const char *valueNames =
        "Default\000Cancel\000Destructive\000";
    static const int32_t values[] = {
        MatchaAlertPBButtonStyle_Default,
        MatchaAlertPBButtonStyle_Cancel,
        MatchaAlertPBButtonStyle_Destructive,
    };
    GPBEnumDescriptor *worker =
        [GPBEnumDescriptor allocDescriptorForName:GPBNSStringifySymbol(MatchaAlertPBButtonStyle)
                                       valueNames:valueNames
                                           values:values
                                            count:(uint32_t)(sizeof(values) / sizeof(int32_t))
                                     enumVerifier:MatchaAlertPBButtonStyle_IsValidValue];
    if (!OSAtomicCompareAndSwapPtrBarrier(nil, worker, (void * volatile *)&descriptor)) {
      [worker release];
    }
  }
  return descriptor;
}

BOOL MatchaAlertPBButtonStyle_IsValidValue(int32_t value__) {
  switch (value__) {
    case MatchaAlertPBButtonStyle_Default:
    case MatchaAlertPBButtonStyle_Cancel:
    case MatchaAlertPBButtonStyle_Destructive:
      return YES;
    default:
      return NO;
  }
}

#pragma mark - MatchaAlertPBView

@implementation MatchaAlertPBView

@dynamic id_p;
@dynamic title;
@dynamic message;
@dynamic buttonsArray, buttonsArray_Count;

typedef struct MatchaAlertPBView__storage_ {
  uint32_t _has_storage_[1];
  NSString *title;
  NSString *message;
  NSMutableArray *buttonsArray;
  int64_t id_p;
} MatchaAlertPBView__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "id_p",
        .dataTypeSpecific.className = NULL,
        .number = MatchaAlertPBView_FieldNumber_Id_p,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaAlertPBView__storage_, id_p),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "title",
        .dataTypeSpecific.className = NULL,
        .number = MatchaAlertPBView_FieldNumber_Title,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(MatchaAlertPBView__storage_, title),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "message",
        .dataTypeSpecific.className = NULL,
        .number = MatchaAlertPBView_FieldNumber_Message,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(MatchaAlertPBView__storage_, message),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "buttonsArray",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaAlertPBButton),
        .number = MatchaAlertPBView_FieldNumber_ButtonsArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(MatchaAlertPBView__storage_, buttonsArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaAlertPBView class]
                                     rootClass:[MatchaAlertPBAlertRoot class]
                                          file:MatchaAlertPBAlertRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaAlertPBView__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaAlertPBButton

@implementation MatchaAlertPBButton

@dynamic title;
@dynamic style;

typedef struct MatchaAlertPBButton__storage_ {
  uint32_t _has_storage_[1];
  MatchaAlertPBButtonStyle style;
  NSString *title;
} MatchaAlertPBButton__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "title",
        .dataTypeSpecific.className = NULL,
        .number = MatchaAlertPBButton_FieldNumber_Title,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaAlertPBButton__storage_, title),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "style",
        .dataTypeSpecific.enumDescFunc = MatchaAlertPBButtonStyle_EnumDescriptor,
        .number = MatchaAlertPBButton_FieldNumber_Style,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(MatchaAlertPBButton__storage_, style),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldHasEnumDescriptor),
        .dataType = GPBDataTypeEnum,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaAlertPBButton class]
                                     rootClass:[MatchaAlertPBAlertRoot class]
                                          file:MatchaAlertPBAlertRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaAlertPBButton__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

int32_t MatchaAlertPBButton_Style_RawValue(MatchaAlertPBButton *message) {
  GPBDescriptor *descriptor = [MatchaAlertPBButton descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:MatchaAlertPBButton_FieldNumber_Style];
  return GPBGetMessageInt32Field(message, field);
}

void SetMatchaAlertPBButton_Style_RawValue(MatchaAlertPBButton *message, int32_t value) {
  GPBDescriptor *descriptor = [MatchaAlertPBButton descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:MatchaAlertPBButton_FieldNumber_Style];
  GPBSetInt32IvarWithFieldInternal(message, field, value, descriptor.file.syntax);
}


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
