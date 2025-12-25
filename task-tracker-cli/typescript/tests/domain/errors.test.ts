/// <reference types="node" />
import { strict as assert } from "node:assert";
import { describe, it } from "node:test";
import { ValidationError } from "../../src/domain/errors.js";

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
});
