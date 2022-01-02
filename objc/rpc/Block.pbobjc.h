// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: rpc/block.proto

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

#if GOOGLE_PROTOBUF_OBJC_VERSION < 30004
#error This file was generated by a newer version of protoc which is incompatible with your Protocol Buffer library sources.
#endif
#if 30004 < GOOGLE_PROTOBUF_OBJC_MIN_SUPPORTED_VERSION
#error This file was generated by an older version of protoc which is incompatible with your Protocol Buffer library sources.
#endif

// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

CF_EXTERN_C_BEGIN

@class Header;
@class ReplyTransaction;
@class ReplyVerify;

NS_ASSUME_NONNULL_BEGIN

#pragma mark - BlockRoot

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
GPB_FINAL @interface BlockRoot : GPBRootObject
@end

#pragma mark - VerifyRequest

typedef GPB_ENUM(VerifyRequest_FieldNumber) {
  VerifyRequest_FieldNumber_Height = 1,
  VerifyRequest_FieldNumber_BlockHash = 2,
  VerifyRequest_FieldNumber_Hash_p = 3,
  VerifyRequest_FieldNumber_Time = 4,
  VerifyRequest_FieldNumber_Nonce = 5,
  VerifyRequest_FieldNumber_Miner = 6,
};

GPB_FINAL @interface VerifyRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSData *height;

@property(nonatomic, readwrite, copy, null_resettable) NSData *blockHash;

@property(nonatomic, readwrite, copy, null_resettable) NSData *hash_p;

@property(nonatomic, readwrite, copy, null_resettable) NSData *time;

@property(nonatomic, readwrite, copy, null_resettable) NSData *nonce;

@property(nonatomic, readwrite, copy, null_resettable) NSData *miner;

@end

#pragma mark - TransactionRequest

typedef GPB_ENUM(TransactionRequest_FieldNumber) {
  TransactionRequest_FieldNumber_Value = 1,
  TransactionRequest_FieldNumber_BaseFee = 2,
  TransactionRequest_FieldNumber_To = 3,
  TransactionRequest_FieldNumber_Random = 4,
  TransactionRequest_FieldNumber_TxHash = 5,
  TransactionRequest_FieldNumber_Time = 6,
  TransactionRequest_FieldNumber_Nonce = 7,
  TransactionRequest_FieldNumber_Sign = 8,
};

GPB_FINAL @interface TransactionRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSData *value;

@property(nonatomic, readwrite, copy, null_resettable) NSData *baseFee;

@property(nonatomic, readwrite, copy, null_resettable) NSData *to;

@property(nonatomic, readwrite, copy, null_resettable) NSData *random;

@property(nonatomic, readwrite, copy, null_resettable) NSData *txHash;

@property(nonatomic, readwrite, copy, null_resettable) NSData *time;

@property(nonatomic, readwrite, copy, null_resettable) NSData *nonce;

@property(nonatomic, readwrite, copy, null_resettable) NSData *sign;

@end

#pragma mark - HashRequest

typedef GPB_ENUM(HashRequest_FieldNumber) {
  HashRequest_FieldNumber_Hash_p = 1,
};

GPB_FINAL @interface HashRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSData *hash_p;

@end

#pragma mark - AddressRequest

typedef GPB_ENUM(AddressRequest_FieldNumber) {
  AddressRequest_FieldNumber_Address = 1,
};

GPB_FINAL @interface AddressRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSData *address;

@end

#pragma mark - ReplyBool

typedef GPB_ENUM(ReplyBool_FieldNumber) {
  ReplyBool_FieldNumber_Ok = 1,
};

GPB_FINAL @interface ReplyBool : GPBMessage

@property(nonatomic, readwrite) BOOL ok;

@end

#pragma mark - ReplyValue

typedef GPB_ENUM(ReplyValue_FieldNumber) {
  ReplyValue_FieldNumber_In_p = 1,
  ReplyValue_FieldNumber_Out_p = 2,
};

GPB_FINAL @interface ReplyValue : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSData *in_p;

@property(nonatomic, readwrite, copy, null_resettable) NSData *out_p;

@end

#pragma mark - ReplyBalance

typedef GPB_ENUM(ReplyBalance_FieldNumber) {
  ReplyBalance_FieldNumber_Balance = 1,
};

GPB_FINAL @interface ReplyBalance : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSData *balance;

@end

#pragma mark - HeightRequest

typedef GPB_ENUM(HeightRequest_FieldNumber) {
  HeightRequest_FieldNumber_Height = 1,
};

GPB_FINAL @interface HeightRequest : GPBMessage

@property(nonatomic, readwrite) int64_t height;

@end

#pragma mark - ReplyBlock

typedef GPB_ENUM(ReplyBlock_FieldNumber) {
  ReplyBlock_FieldNumber_Height = 1,
  ReplyBlock_FieldNumber_Header = 2,
  ReplyBlock_FieldNumber_TransactionArray = 3,
};

GPB_FINAL @interface ReplyBlock : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSData *height;

@property(nonatomic, readwrite, strong, null_resettable) Header *header;
/** Test to see if @c header has been set. */
@property(nonatomic, readwrite) BOOL hasHeader;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<ReplyTransaction*> *transactionArray;
/** The number of items in @c transactionArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger transactionArray_Count;

@end

#pragma mark - ReplyLatest

typedef GPB_ENUM(ReplyLatest_FieldNumber) {
  ReplyLatest_FieldNumber_Height = 1,
  ReplyLatest_FieldNumber_Hash_p = 2,
};

GPB_FINAL @interface ReplyLatest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSData *height;

@property(nonatomic, readwrite, copy, null_resettable) NSData *hash_p;

@end

#pragma mark - Header

typedef GPB_ENUM(Header_FieldNumber) {
  Header_FieldNumber_ParentHash = 1,
  Header_FieldNumber_Hash_p = 2,
  Header_FieldNumber_Time = 3,
  Header_FieldNumber_Nonce = 4,
  Header_FieldNumber_Miner = 5,
};

GPB_FINAL @interface Header : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSData *parentHash;

@property(nonatomic, readwrite, copy, null_resettable) NSData *hash_p;

@property(nonatomic, readwrite, copy, null_resettable) NSData *time;

@property(nonatomic, readwrite, copy, null_resettable) NSData *nonce;

@property(nonatomic, readwrite, copy, null_resettable) NSData *miner;

@end

#pragma mark - ReplyTransaction

typedef GPB_ENUM(ReplyTransaction_FieldNumber) {
  ReplyTransaction_FieldNumber_Value = 1,
  ReplyTransaction_FieldNumber_BaseFee = 2,
  ReplyTransaction_FieldNumber_To = 3,
  ReplyTransaction_FieldNumber_Random = 4,
  ReplyTransaction_FieldNumber_TxHash = 5,
  ReplyTransaction_FieldNumber_Time = 6,
  ReplyTransaction_FieldNumber_Nonce = 7,
  ReplyTransaction_FieldNumber_Sign = 8,
  ReplyTransaction_FieldNumber_State = 9,
};

GPB_FINAL @interface ReplyTransaction : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSData *value;

@property(nonatomic, readwrite, copy, null_resettable) NSData *baseFee;

@property(nonatomic, readwrite, copy, null_resettable) NSData *to;

@property(nonatomic, readwrite, copy, null_resettable) NSData *random;

@property(nonatomic, readwrite, copy, null_resettable) NSData *txHash;

@property(nonatomic, readwrite, copy, null_resettable) NSData *time;

@property(nonatomic, readwrite, copy, null_resettable) NSData *nonce;

@property(nonatomic, readwrite, copy, null_resettable) NSData *sign;

@property(nonatomic, readwrite) BOOL state;

@end

#pragma mark - ReplyVerifys

typedef GPB_ENUM(ReplyVerifys_FieldNumber) {
  ReplyVerifys_FieldNumber_VerifysArray = 1,
};

GPB_FINAL @interface ReplyVerifys : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<ReplyVerify*> *verifysArray;
/** The number of items in @c verifysArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger verifysArray_Count;

@end

#pragma mark - ReplyVerify

typedef GPB_ENUM(ReplyVerify_FieldNumber) {
  ReplyVerify_FieldNumber_Height = 1,
  ReplyVerify_FieldNumber_BlockHash = 2,
  ReplyVerify_FieldNumber_Hash_p = 3,
  ReplyVerify_FieldNumber_Time = 4,
  ReplyVerify_FieldNumber_Nonce = 5,
  ReplyVerify_FieldNumber_Miner = 6,
};

GPB_FINAL @interface ReplyVerify : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSData *height;

@property(nonatomic, readwrite, copy, null_resettable) NSData *blockHash;

@property(nonatomic, readwrite, copy, null_resettable) NSData *hash_p;

@property(nonatomic, readwrite, copy, null_resettable) NSData *time;

@property(nonatomic, readwrite, copy, null_resettable) NSData *nonce;

@property(nonatomic, readwrite, copy, null_resettable) NSData *miner;

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
