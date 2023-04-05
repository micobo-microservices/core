// @generated by protoc-gen-es v1.2.0
// @generated from file webhook/v1/webhook.proto (package webhook.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { IDnowAbortedPayload, IDnowCanceledPayload, IDnowFraudSuspicionConfirmedPayload, IDnowFraudSuspicionPendingPayload, IDnowReviewPendingPayload, IDnowSuccessDataChangedPayload, IDnowSuccessPayload } from "../../provider/v1/idnow_pb.js";

/**
 * @generated from enum webhook.v1.HookTopic
 */
export const HookTopic = proto3.makeEnum(
  "webhook.v1.HookTopic",
  [
    {no: 0, name: "HOOK_TOPIC_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "HOOK_TOPIC_IDNOW_SUCCESS", localName: "IDNOW_SUCCESS"},
    {no: 2, name: "HOOK_TOPIC_IDNOW_REVIEW_PENDING", localName: "IDNOW_REVIEW_PENDING"},
    {no: 3, name: "HOOK_TOPIC_IDNOW_FRAUD_SUSPICION_PENDING", localName: "IDNOW_FRAUD_SUSPICION_PENDING"},
    {no: 4, name: "HOOK_TOPIC_IDNOW_SUCCESS_DATA_CHANGED", localName: "IDNOW_SUCCESS_DATA_CHANGED"},
    {no: 5, name: "HOOK_TOPIC_IDNOW_FRAUD_SUSPICION_CONFIRMED", localName: "IDNOW_FRAUD_SUSPICION_CONFIRMED"},
    {no: 6, name: "HOOK_TOPIC_IDNOW_CANCELED", localName: "IDNOW_CANCELED"},
    {no: 7, name: "HOOK_TOPIC_IDNOW_ABORTED", localName: "IDNOW_ABORTED"},
  ],
);

/**
 * @generated from message webhook.v1.IDnowAbortedRequest
 */
export const IDnowAbortedRequest = proto3.makeMessageType(
  "webhook.v1.IDnowAbortedRequest",
  () => [
    { no: 1, name: "tenant_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "request", kind: "message", T: IDnowAbortedPayload },
  ],
);

/**
 * @generated from message webhook.v1.IDnowCanceledRequest
 */
export const IDnowCanceledRequest = proto3.makeMessageType(
  "webhook.v1.IDnowCanceledRequest",
  () => [
    { no: 1, name: "tenant_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "request", kind: "message", T: IDnowCanceledPayload },
  ],
);

/**
 * @generated from message webhook.v1.IDnowFraudSuspicionConfirmedRequest
 */
export const IDnowFraudSuspicionConfirmedRequest = proto3.makeMessageType(
  "webhook.v1.IDnowFraudSuspicionConfirmedRequest",
  () => [
    { no: 1, name: "tenant_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "request", kind: "message", T: IDnowFraudSuspicionConfirmedPayload },
  ],
);

/**
 * @generated from message webhook.v1.IDnowSuccessDataChangedRequest
 */
export const IDnowSuccessDataChangedRequest = proto3.makeMessageType(
  "webhook.v1.IDnowSuccessDataChangedRequest",
  () => [
    { no: 1, name: "tenant_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "request", kind: "message", T: IDnowSuccessDataChangedPayload },
  ],
);

/**
 * @generated from message webhook.v1.IDnowFraudSuspicionPendingRequest
 */
export const IDnowFraudSuspicionPendingRequest = proto3.makeMessageType(
  "webhook.v1.IDnowFraudSuspicionPendingRequest",
  () => [
    { no: 1, name: "tenant_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "request", kind: "message", T: IDnowFraudSuspicionPendingPayload },
  ],
);

/**
 * @generated from message webhook.v1.IDnowReviewPendingRequest
 */
export const IDnowReviewPendingRequest = proto3.makeMessageType(
  "webhook.v1.IDnowReviewPendingRequest",
  () => [
    { no: 1, name: "tenant_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "request", kind: "message", T: IDnowReviewPendingPayload },
  ],
);

/**
 * @generated from message webhook.v1.IDnowSuccessRequest
 */
export const IDnowSuccessRequest = proto3.makeMessageType(
  "webhook.v1.IDnowSuccessRequest",
  () => [
    { no: 1, name: "tenant_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "request", kind: "message", T: IDnowSuccessPayload },
  ],
);

