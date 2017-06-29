// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: github.com/overcyn/matcha/pb/view/slider/slider.proto

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

 #import "github.com/gomatcha/matcha/pb/view/slider/Slider.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

#pragma mark - MatchaSliderPbSliderRoot

@implementation MatchaSliderPbSliderRoot

// No extensions in the file and no imports, so no need to generate
// +extensionRegistry.

@end

#pragma mark - MatchaSliderPbSliderRoot_FileDescriptor

static GPBFileDescriptor *MatchaSliderPbSliderRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@"slider"
                                                 objcPrefix:@"MatchaSliderPb"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - MatchaSliderPbView

@implementation MatchaSliderPbView

@dynamic value;
@dynamic maxValue;
@dynamic minValue;
@dynamic enabled;

typedef struct MatchaSliderPbView__storage_ {
  uint32_t _has_storage_[1];
  double value;
  double maxValue;
  double minValue;
} MatchaSliderPbView__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "value",
        .dataTypeSpecific.className = NULL,
        .number = MatchaSliderPbView_FieldNumber_Value,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaSliderPbView__storage_, value),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeDouble,
      },
      {
        .name = "maxValue",
        .dataTypeSpecific.className = NULL,
        .number = MatchaSliderPbView_FieldNumber_MaxValue,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(MatchaSliderPbView__storage_, maxValue),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeDouble,
      },
      {
        .name = "minValue",
        .dataTypeSpecific.className = NULL,
        .number = MatchaSliderPbView_FieldNumber_MinValue,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(MatchaSliderPbView__storage_, minValue),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeDouble,
      },
      {
        .name = "enabled",
        .dataTypeSpecific.className = NULL,
        .number = MatchaSliderPbView_FieldNumber_Enabled,
        .hasIndex = 3,
        .offset = 4,  // Stored in _has_storage_ to save space.
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeBool,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaSliderPbView class]
                                     rootClass:[MatchaSliderPbSliderRoot class]
                                          file:MatchaSliderPbSliderRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaSliderPbView__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\002\002\010\000\003\010\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaSliderPbEvent

@implementation MatchaSliderPbEvent

@dynamic value;

typedef struct MatchaSliderPbEvent__storage_ {
  uint32_t _has_storage_[1];
  double value;
} MatchaSliderPbEvent__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "value",
        .dataTypeSpecific.className = NULL,
        .number = MatchaSliderPbEvent_FieldNumber_Value,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaSliderPbEvent__storage_, value),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeDouble,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaSliderPbEvent class]
                                     rootClass:[MatchaSliderPbSliderRoot class]
                                          file:MatchaSliderPbSliderRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaSliderPbEvent__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
