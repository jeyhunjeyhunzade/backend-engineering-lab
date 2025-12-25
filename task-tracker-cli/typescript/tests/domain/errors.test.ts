/// <reference types="node" />
import { strict as assert } from "node:assert";
import { describe, it } from "node:test";
import {
  ValidationError,
  NotFoundError,
  StorageError,
} from "../../src/domain/errors.js";

describe("ValidationError", () => {
  it("should be an instance of Error", () => {
    const error = new ValidationError("Test error");
    assert.ok(error instanceof Error);
    assert.ok(error instanceof ValidationError);
  });

  it("should have the correct message", () => {
    const message = "This is a validation error";
    const error = new ValidationError(message);
    assert.equal(error.message, message);
  });

  it("should have the correct name", () => {
    const error = new ValidationError("Test error");
    assert.equal(error.name, "ValidationError");
  });

  it("should support optional cause parameter", () => {
    const cause = new Error("Original error");
    const error = new ValidationError("Test error", cause);
    assert.equal(error.cause, cause);
  });
});

describe("NotFoundError", () => {
  it("should be an instance of Error", () => {
    const error = new NotFoundError("Test error");
    assert.ok(error instanceof Error);
    assert.ok(error instanceof NotFoundError);
  });

  it("should have the correct message", () => {
    const message = "Resource not found";
    const error = new NotFoundError(message);
    assert.equal(error.message, message);
  });

  it("should have the correct name", () => {
    const error = new NotFoundError("Test error");
    assert.equal(error.name, "NotFoundError");
  });

  it("should support optional cause parameter", () => {
    const cause = new Error("Original error");
    const error = new NotFoundError("Test error", cause);
    assert.equal(error.cause, cause);
  });
});

describe("StorageError", () => {
  it("should be an instance of Error", () => {
    const error = new StorageError("Test error");
    assert.ok(error instanceof Error);
    assert.ok(error instanceof StorageError);
  });

  it("should have the correct message", () => {
    const message = "Storage operation failed";
    const error = new StorageError(message);
    assert.equal(error.message, message);
  });

  it("should have the correct name", () => {
    const error = new StorageError("Test error");
    assert.equal(error.name, "StorageError");
  });

  it("should support optional cause parameter", () => {
    const cause = new Error("Original error");
    const error = new StorageError("Test error", cause);
    assert.equal(error.cause, cause);
  });
});
