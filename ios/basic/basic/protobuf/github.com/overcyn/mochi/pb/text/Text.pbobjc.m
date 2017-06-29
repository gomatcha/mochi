// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: github.com/overcyn/matcha/pb/text/text.proto

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

 #import "github.com/gomatcha/matcha/pb/text/Text.pbobjc.h"
 #import "github.com/gomatcha/matcha/pb/Layout.pbobjc.h"
 #import "github.com/gomatcha/matcha/pb/Color.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

#pragma mark - MatchaPBTextRoot

@implementation MatchaPBTextRoot

// No extensions in the file and none of the imports (direct or indirect)
// defined extensions, so no need to generate +extensionRegistry.

@end

#pragma mark - MatchaPBTextRoot_FileDescriptor

static GPBFileDescriptor *MatchaPBTextRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@"matcha.text"
                                                 objcPrefix:@"MatchaPB"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - Enum MatchaPBTextAlignment

GPBEnumDescriptor *MatchaPBTextAlignment_EnumDescriptor(void) {
  static GPBEnumDescriptor *descriptor = NULL;
  if (!descriptor) {
    static const char *valueNames =
        "TextAlignmentLeft\000TextAlignmentRight\000Tex"
        "tAlignmentCenter\000TextAlignmentJustified\000";
    static const int32_t values[] = {
        MatchaPBTextAlignment_TextAlignmentLeft,
        MatchaPBTextAlignment_TextAlignmentRight,
        MatchaPBTextAlignment_TextAlignmentCenter,
        MatchaPBTextAlignment_TextAlignmentJustified,
    };
    GPBEnumDescriptor *worker =
        [GPBEnumDescriptor allocDescriptorForName:GPBNSStringifySymbol(MatchaPBTextAlignment)
                                       valueNames:valueNames
                                           values:values
                                            count:(uint32_t)(sizeof(values) / sizeof(int32_t))
                                     enumVerifier:MatchaPBTextAlignment_IsValidValue];
    if (!OSAtomicCompareAndSwapPtrBarrier(nil, worker, (void * volatile *)&descriptor)) {
      [worker release];
    }
  }
  return descriptor;
}

BOOL MatchaPBTextAlignment_IsValidValue(int32_t value__) {
  switch (value__) {
    case MatchaPBTextAlignment_TextAlignmentLeft:
    case MatchaPBTextAlignment_TextAlignmentRight:
    case MatchaPBTextAlignment_TextAlignmentCenter:
    case MatchaPBTextAlignment_TextAlignmentJustified:
      return YES;
    default:
      return NO;
  }
}

#pragma mark - Enum MatchaPBStrikethroughStyle

GPBEnumDescriptor *MatchaPBStrikethroughStyle_EnumDescriptor(void) {
  static GPBEnumDescriptor *descriptor = NULL;
  if (!descriptor) {
    static const char *valueNames =
        "StrikethroughStyleNone\000StrikethroughStyl"
        "eSingle\000StrikethroughStyleDouble\000Striket"
        "hroughStyleThick\000StrikethroughStyleDotte"
        "d\000StrikethroughStyleDashed\000";
    static const int32_t values[] = {
        MatchaPBStrikethroughStyle_StrikethroughStyleNone,
        MatchaPBStrikethroughStyle_StrikethroughStyleSingle,
        MatchaPBStrikethroughStyle_StrikethroughStyleDouble,
        MatchaPBStrikethroughStyle_StrikethroughStyleThick,
        MatchaPBStrikethroughStyle_StrikethroughStyleDotted,
        MatchaPBStrikethroughStyle_StrikethroughStyleDashed,
    };
    GPBEnumDescriptor *worker =
        [GPBEnumDescriptor allocDescriptorForName:GPBNSStringifySymbol(MatchaPBStrikethroughStyle)
                                       valueNames:valueNames
                                           values:values
                                            count:(uint32_t)(sizeof(values) / sizeof(int32_t))
                                     enumVerifier:MatchaPBStrikethroughStyle_IsValidValue];
    if (!OSAtomicCompareAndSwapPtrBarrier(nil, worker, (void * volatile *)&descriptor)) {
      [worker release];
    }
  }
  return descriptor;
}

BOOL MatchaPBStrikethroughStyle_IsValidValue(int32_t value__) {
  switch (value__) {
    case MatchaPBStrikethroughStyle_StrikethroughStyleNone:
    case MatchaPBStrikethroughStyle_StrikethroughStyleSingle:
    case MatchaPBStrikethroughStyle_StrikethroughStyleDouble:
    case MatchaPBStrikethroughStyle_StrikethroughStyleThick:
    case MatchaPBStrikethroughStyle_StrikethroughStyleDotted:
    case MatchaPBStrikethroughStyle_StrikethroughStyleDashed:
      return YES;
    default:
      return NO;
  }
}

#pragma mark - Enum MatchaPBUnderlineStyle

GPBEnumDescriptor *MatchaPBUnderlineStyle_EnumDescriptor(void) {
  static GPBEnumDescriptor *descriptor = NULL;
  if (!descriptor) {
    static const char *valueNames =
        "UndrelineStyleNone\000UndrelineStyleSingle\000"
        "UndrelineStyleDouble\000UndrelineStyleThick"
        "\000UndrelineStyleDotted\000UndrelineStyleDash"
        "ed\000";
    static const int32_t values[] = {
        MatchaPBUnderlineStyle_UndrelineStyleNone,
        MatchaPBUnderlineStyle_UndrelineStyleSingle,
        MatchaPBUnderlineStyle_UndrelineStyleDouble,
        MatchaPBUnderlineStyle_UndrelineStyleThick,
        MatchaPBUnderlineStyle_UndrelineStyleDotted,
        MatchaPBUnderlineStyle_UndrelineStyleDashed,
    };
    GPBEnumDescriptor *worker =
        [GPBEnumDescriptor allocDescriptorForName:GPBNSStringifySymbol(MatchaPBUnderlineStyle)
                                       valueNames:valueNames
                                           values:values
                                            count:(uint32_t)(sizeof(values) / sizeof(int32_t))
                                     enumVerifier:MatchaPBUnderlineStyle_IsValidValue];
    if (!OSAtomicCompareAndSwapPtrBarrier(nil, worker, (void * volatile *)&descriptor)) {
      [worker release];
    }
  }
  return descriptor;
}

BOOL MatchaPBUnderlineStyle_IsValidValue(int32_t value__) {
  switch (value__) {
    case MatchaPBUnderlineStyle_UndrelineStyleNone:
    case MatchaPBUnderlineStyle_UndrelineStyleSingle:
    case MatchaPBUnderlineStyle_UndrelineStyleDouble:
    case MatchaPBUnderlineStyle_UndrelineStyleThick:
    case MatchaPBUnderlineStyle_UndrelineStyleDotted:
    case MatchaPBUnderlineStyle_UndrelineStyleDashed:
      return YES;
    default:
      return NO;
  }
}

#pragma mark - Enum MatchaPBTextWrap

GPBEnumDescriptor *MatchaPBTextWrap_EnumDescriptor(void) {
  static GPBEnumDescriptor *descriptor = NULL;
  if (!descriptor) {
    static const char *valueNames =
        "TextWrapNone\000TextWrapWord\000TextWrapCharac"
        "ter\000";
    static const int32_t values[] = {
        MatchaPBTextWrap_TextWrapNone,
        MatchaPBTextWrap_TextWrapWord,
        MatchaPBTextWrap_TextWrapCharacter,
    };
    GPBEnumDescriptor *worker =
        [GPBEnumDescriptor allocDescriptorForName:GPBNSStringifySymbol(MatchaPBTextWrap)
                                       valueNames:valueNames
                                           values:values
                                            count:(uint32_t)(sizeof(values) / sizeof(int32_t))
                                     enumVerifier:MatchaPBTextWrap_IsValidValue];
    if (!OSAtomicCompareAndSwapPtrBarrier(nil, worker, (void * volatile *)&descriptor)) {
      [worker release];
    }
  }
  return descriptor;
}

BOOL MatchaPBTextWrap_IsValidValue(int32_t value__) {
  switch (value__) {
    case MatchaPBTextWrap_TextWrapNone:
    case MatchaPBTextWrap_TextWrapWord:
    case MatchaPBTextWrap_TextWrapCharacter:
      return YES;
    default:
      return NO;
  }
}

#pragma mark - Enum MatchaPBTruncation

GPBEnumDescriptor *MatchaPBTruncation_EnumDescriptor(void) {
  static GPBEnumDescriptor *descriptor = NULL;
  if (!descriptor) {
    static const char *valueNames =
        "TruncationNone\000TruncationStart\000Truncatio"
        "nMiddle\000TruncationEnd\000";
    static const int32_t values[] = {
        MatchaPBTruncation_TruncationNone,
        MatchaPBTruncation_TruncationStart,
        MatchaPBTruncation_TruncationMiddle,
        MatchaPBTruncation_TruncationEnd,
    };
    GPBEnumDescriptor *worker =
        [GPBEnumDescriptor allocDescriptorForName:GPBNSStringifySymbol(MatchaPBTruncation)
                                       valueNames:valueNames
                                           values:values
                                            count:(uint32_t)(sizeof(values) / sizeof(int32_t))
                                     enumVerifier:MatchaPBTruncation_IsValidValue];
    if (!OSAtomicCompareAndSwapPtrBarrier(nil, worker, (void * volatile *)&descriptor)) {
      [worker release];
    }
  }
  return descriptor;
}

BOOL MatchaPBTruncation_IsValidValue(int32_t value__) {
  switch (value__) {
    case MatchaPBTruncation_TruncationNone:
    case MatchaPBTruncation_TruncationStart:
    case MatchaPBTruncation_TruncationMiddle:
    case MatchaPBTruncation_TruncationEnd:
      return YES;
    default:
      return NO;
  }
}

#pragma mark - MatchaPBSizeFunc

@implementation MatchaPBSizeFunc

@dynamic hasText, text;
@dynamic hasMinSize, minSize;
@dynamic hasMaxSize, maxSize;

typedef struct MatchaPBSizeFunc__storage_ {
  uint32_t _has_storage_[1];
  MatchaPBStyledText *text;
  MatchaPBPoint *minSize;
  MatchaPBPoint *maxSize;
} MatchaPBSizeFunc__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "text",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBStyledText),
        .number = MatchaPBSizeFunc_FieldNumber_Text,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaPBSizeFunc__storage_, text),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "minSize",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBPoint),
        .number = MatchaPBSizeFunc_FieldNumber_MinSize,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(MatchaPBSizeFunc__storage_, minSize),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "maxSize",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBPoint),
        .number = MatchaPBSizeFunc_FieldNumber_MaxSize,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(MatchaPBSizeFunc__storage_, maxSize),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaPBSizeFunc class]
                                     rootClass:[MatchaPBTextRoot class]
                                          file:MatchaPBTextRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaPBSizeFunc__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\002\002\007\000\003\007\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaPBText

@implementation MatchaPBText

@dynamic text;

typedef struct MatchaPBText__storage_ {
  uint32_t _has_storage_[1];
  NSString *text;
} MatchaPBText__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "text",
        .dataTypeSpecific.className = NULL,
        .number = MatchaPBText_FieldNumber_Text,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaPBText__storage_, text),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaPBText class]
                                     rootClass:[MatchaPBTextRoot class]
                                          file:MatchaPBTextRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaPBText__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaPBStyledText

@implementation MatchaPBStyledText

@dynamic hasStyle, style;
@dynamic hasText, text;

typedef struct MatchaPBStyledText__storage_ {
  uint32_t _has_storage_[1];
  MatchaPBTextStyle *style;
  MatchaPBText *text;
} MatchaPBStyledText__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "style",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBTextStyle),
        .number = MatchaPBStyledText_FieldNumber_Style,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaPBStyledText__storage_, style),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "text",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBText),
        .number = MatchaPBStyledText_FieldNumber_Text,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(MatchaPBStyledText__storage_, text),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaPBStyledText class]
                                     rootClass:[MatchaPBTextRoot class]
                                          file:MatchaPBTextRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaPBStyledText__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaPBFont

@implementation MatchaPBFont

@dynamic family;
@dynamic face;
@dynamic size;

typedef struct MatchaPBFont__storage_ {
  uint32_t _has_storage_[1];
  NSString *family;
  NSString *face;
  double size;
} MatchaPBFont__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "family",
        .dataTypeSpecific.className = NULL,
        .number = MatchaPBFont_FieldNumber_Family,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaPBFont__storage_, family),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "face",
        .dataTypeSpecific.className = NULL,
        .number = MatchaPBFont_FieldNumber_Face,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(MatchaPBFont__storage_, face),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "size",
        .dataTypeSpecific.className = NULL,
        .number = MatchaPBFont_FieldNumber_Size,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(MatchaPBFont__storage_, size),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeDouble,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaPBFont class]
                                     rootClass:[MatchaPBTextRoot class]
                                          file:MatchaPBTextRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaPBFont__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaPBTextStyle

@implementation MatchaPBTextStyle

@dynamic textAlignment;
@dynamic strikethroughStyle;
@dynamic hasStrikethroughColor, strikethroughColor;
@dynamic underlineStyle;
@dynamic hasUnderlineColor, underlineColor;
@dynamic hasFont, font;
@dynamic hyphenation;
@dynamic lineHeightMultiple;
@dynamic maxLines;
@dynamic hasTextColor, textColor;
@dynamic wrap;
@dynamic truncation;
@dynamic truncationString;

typedef struct MatchaPBTextStyle__storage_ {
  uint32_t _has_storage_[1];
  MatchaPBTextAlignment textAlignment;
  MatchaPBStrikethroughStyle strikethroughStyle;
  MatchaPBUnderlineStyle underlineStyle;
  MatchaPBTextWrap wrap;
  MatchaPBTruncation truncation;
  MatchaPBColor *strikethroughColor;
  MatchaPBColor *underlineColor;
  MatchaPBFont *font;
  MatchaPBColor *textColor;
  NSString *truncationString;
  double hyphenation;
  double lineHeightMultiple;
  int64_t maxLines;
} MatchaPBTextStyle__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "textAlignment",
        .dataTypeSpecific.enumDescFunc = MatchaPBTextAlignment_EnumDescriptor,
        .number = MatchaPBTextStyle_FieldNumber_TextAlignment,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaPBTextStyle__storage_, textAlignment),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom | GPBFieldHasEnumDescriptor),
        .dataType = GPBDataTypeEnum,
      },
      {
        .name = "strikethroughStyle",
        .dataTypeSpecific.enumDescFunc = MatchaPBStrikethroughStyle_EnumDescriptor,
        .number = MatchaPBTextStyle_FieldNumber_StrikethroughStyle,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(MatchaPBTextStyle__storage_, strikethroughStyle),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom | GPBFieldHasEnumDescriptor),
        .dataType = GPBDataTypeEnum,
      },
      {
        .name = "strikethroughColor",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBColor),
        .number = MatchaPBTextStyle_FieldNumber_StrikethroughColor,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(MatchaPBTextStyle__storage_, strikethroughColor),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "underlineStyle",
        .dataTypeSpecific.enumDescFunc = MatchaPBUnderlineStyle_EnumDescriptor,
        .number = MatchaPBTextStyle_FieldNumber_UnderlineStyle,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(MatchaPBTextStyle__storage_, underlineStyle),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom | GPBFieldHasEnumDescriptor),
        .dataType = GPBDataTypeEnum,
      },
      {
        .name = "underlineColor",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBColor),
        .number = MatchaPBTextStyle_FieldNumber_UnderlineColor,
        .hasIndex = 4,
        .offset = (uint32_t)offsetof(MatchaPBTextStyle__storage_, underlineColor),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "font",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBFont),
        .number = MatchaPBTextStyle_FieldNumber_Font,
        .hasIndex = 5,
        .offset = (uint32_t)offsetof(MatchaPBTextStyle__storage_, font),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "hyphenation",
        .dataTypeSpecific.className = NULL,
        .number = MatchaPBTextStyle_FieldNumber_Hyphenation,
        .hasIndex = 6,
        .offset = (uint32_t)offsetof(MatchaPBTextStyle__storage_, hyphenation),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeDouble,
      },
      {
        .name = "lineHeightMultiple",
        .dataTypeSpecific.className = NULL,
        .number = MatchaPBTextStyle_FieldNumber_LineHeightMultiple,
        .hasIndex = 7,
        .offset = (uint32_t)offsetof(MatchaPBTextStyle__storage_, lineHeightMultiple),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeDouble,
      },
      {
        .name = "maxLines",
        .dataTypeSpecific.className = NULL,
        .number = MatchaPBTextStyle_FieldNumber_MaxLines,
        .hasIndex = 8,
        .offset = (uint32_t)offsetof(MatchaPBTextStyle__storage_, maxLines),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "textColor",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBColor),
        .number = MatchaPBTextStyle_FieldNumber_TextColor,
        .hasIndex = 9,
        .offset = (uint32_t)offsetof(MatchaPBTextStyle__storage_, textColor),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "wrap",
        .dataTypeSpecific.enumDescFunc = MatchaPBTextWrap_EnumDescriptor,
        .number = MatchaPBTextStyle_FieldNumber_Wrap,
        .hasIndex = 10,
        .offset = (uint32_t)offsetof(MatchaPBTextStyle__storage_, wrap),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldHasEnumDescriptor),
        .dataType = GPBDataTypeEnum,
      },
      {
        .name = "truncation",
        .dataTypeSpecific.enumDescFunc = MatchaPBTruncation_EnumDescriptor,
        .number = MatchaPBTextStyle_FieldNumber_Truncation,
        .hasIndex = 11,
        .offset = (uint32_t)offsetof(MatchaPBTextStyle__storage_, truncation),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldHasEnumDescriptor),
        .dataType = GPBDataTypeEnum,
      },
      {
        .name = "truncationString",
        .dataTypeSpecific.className = NULL,
        .number = MatchaPBTextStyle_FieldNumber_TruncationString,
        .hasIndex = 12,
        .offset = (uint32_t)offsetof(MatchaPBTextStyle__storage_, truncationString),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaPBTextStyle class]
                                     rootClass:[MatchaPBTextRoot class]
                                          file:MatchaPBTextRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaPBTextStyle__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\t\002\r\000\004\022\000\006\022\000\010\016\000\n\016\000\020\022\000\022\010\000\024\t\000\032\020\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

int32_t MatchaPBTextStyle_TextAlignment_RawValue(MatchaPBTextStyle *message) {
  GPBDescriptor *descriptor = [MatchaPBTextStyle descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:MatchaPBTextStyle_FieldNumber_TextAlignment];
  return GPBGetMessageInt32Field(message, field);
}

void SetMatchaPBTextStyle_TextAlignment_RawValue(MatchaPBTextStyle *message, int32_t value) {
  GPBDescriptor *descriptor = [MatchaPBTextStyle descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:MatchaPBTextStyle_FieldNumber_TextAlignment];
  GPBSetInt32IvarWithFieldInternal(message, field, value, descriptor.file.syntax);
}

int32_t MatchaPBTextStyle_StrikethroughStyle_RawValue(MatchaPBTextStyle *message) {
  GPBDescriptor *descriptor = [MatchaPBTextStyle descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:MatchaPBTextStyle_FieldNumber_StrikethroughStyle];
  return GPBGetMessageInt32Field(message, field);
}

void SetMatchaPBTextStyle_StrikethroughStyle_RawValue(MatchaPBTextStyle *message, int32_t value) {
  GPBDescriptor *descriptor = [MatchaPBTextStyle descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:MatchaPBTextStyle_FieldNumber_StrikethroughStyle];
  GPBSetInt32IvarWithFieldInternal(message, field, value, descriptor.file.syntax);
}

int32_t MatchaPBTextStyle_UnderlineStyle_RawValue(MatchaPBTextStyle *message) {
  GPBDescriptor *descriptor = [MatchaPBTextStyle descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:MatchaPBTextStyle_FieldNumber_UnderlineStyle];
  return GPBGetMessageInt32Field(message, field);
}

void SetMatchaPBTextStyle_UnderlineStyle_RawValue(MatchaPBTextStyle *message, int32_t value) {
  GPBDescriptor *descriptor = [MatchaPBTextStyle descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:MatchaPBTextStyle_FieldNumber_UnderlineStyle];
  GPBSetInt32IvarWithFieldInternal(message, field, value, descriptor.file.syntax);
}

int32_t MatchaPBTextStyle_Wrap_RawValue(MatchaPBTextStyle *message) {
  GPBDescriptor *descriptor = [MatchaPBTextStyle descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:MatchaPBTextStyle_FieldNumber_Wrap];
  return GPBGetMessageInt32Field(message, field);
}

void SetMatchaPBTextStyle_Wrap_RawValue(MatchaPBTextStyle *message, int32_t value) {
  GPBDescriptor *descriptor = [MatchaPBTextStyle descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:MatchaPBTextStyle_FieldNumber_Wrap];
  GPBSetInt32IvarWithFieldInternal(message, field, value, descriptor.file.syntax);
}

int32_t MatchaPBTextStyle_Truncation_RawValue(MatchaPBTextStyle *message) {
  GPBDescriptor *descriptor = [MatchaPBTextStyle descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:MatchaPBTextStyle_FieldNumber_Truncation];
  return GPBGetMessageInt32Field(message, field);
}

void SetMatchaPBTextStyle_Truncation_RawValue(MatchaPBTextStyle *message, int32_t value) {
  GPBDescriptor *descriptor = [MatchaPBTextStyle descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:MatchaPBTextStyle_FieldNumber_Truncation];
  GPBSetInt32IvarWithFieldInternal(message, field, value, descriptor.file.syntax);
}


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
