/// <reference types="node" />
import { strict as assert } from "node:assert";
import { describe, it } from "node:test";
import { isValidStatus, type Task } from "../../src/domain/task.js";
import { nowIso } from "../../src/utils/time.js";

describe("Task Domain", () => {
  describe("isValidStatus", () => {
    it("returns true for valid todo status", () => {
      assert.ok(isValidStatus("todo"));
    });

    it("returns true for valid in-progress status", () => {
      assert.ok(isValidStatus("in-progress"));
    });

    it("returns true for valid done status", () => {
      assert.ok(isValidStatus("done"));
    });

    it("returns false for invalid status", () => {
      assert.equal(isValidStatus("invalid"), false);
      assert.equal(isValidStatus(""), false);
      assert.equal(isValidStatus("DONE"), false);
      assert.equal(isValidStatus("in progress"), false);
    });
  });

  describe("Task creation", () => {
    it("creates a task with all required fields", () => {
      const task: Task = {
        id: 1,
        description: "Test task",
        status: "todo",
        createdAt: nowIso(),
        updatedAt: nowIso(),
      };

      assert.equal(task.id, 1);
      assert.equal(task.description, "Test task");
      assert.equal(task.status, "todo");
      assert.ok(task.createdAt);
      assert.ok(task.updatedAt);
    });

    it("ensures status is one of the valid types", () => {
      const validStatuses: Task["status"][] = ["todo", "in-progress", "done"];

      validStatuses.forEach((status) => {
        const task: Task = {
          id: 1,
          description: "Test task",
          status: status,
          createdAt: nowIso(),
          updatedAt: nowIso(),
        };

        assert.ok(isValidStatus(task.status));
      });
    });

    it("timestamps should be valid ISO strings", () => {
      const task: Task = {
        id: 1,
        description: "Test task",
        status: "todo",
        createdAt: nowIso(),
        updatedAt: nowIso(),
      };

      // Should be parseable as dates
      assert.ok(!isNaN(Date.parse(task.createdAt)));
      assert.ok(!isNaN(Date.parse(task.updatedAt)));
    });
  });

  describe("Task state transitions", () => {
    it("task can transition from todo to in-progress", () => {
      const task: Task = {
        id: 1,
        description: "Test task",
        status: "todo",
        createdAt: nowIso(),
        updatedAt: nowIso(),
      };

      task.status = "in-progress";
      assert.equal(task.status, "in-progress");
    });

    it("task can transition from in-progress to done", () => {
      const task: Task = {
        id: 1,
        description: "Test task",
        status: "in-progress",
        createdAt: nowIso(),
        updatedAt: nowIso(),
      };

      task.status = "done";
      assert.equal(task.status, "done");
    });

    it("task can transition from todo directly to done", () => {
      const task: Task = {
        id: 1,
        description: "Test task",
        status: "todo",
        createdAt: nowIso(),
        updatedAt: nowIso(),
      };

      task.status = "done";
      assert.equal(task.status, "done");
    });
  });

  describe("Task immutability patterns", () => {
    it("task description can be updated", async () => {
      const task: Task = {
        id: 1,
        description: "Original description",
        status: "todo",
        createdAt: nowIso(),
        updatedAt: nowIso(),
      };

      const originalUpdatedAt = task.updatedAt;
      await new Promise((r) => setTimeout(r, 10)); // Ensure timestamp difference
      task.description = "Updated description";
      task.updatedAt = nowIso();

      assert.equal(task.description, "Updated description");
      assert.notEqual(task.updatedAt, originalUpdatedAt);
    });

    it("task ID should remain constant", () => {
      const task: Task = {
        id: 1,
        description: "Test task",
        status: "todo",
        createdAt: nowIso(),
        updatedAt: nowIso(),
      };

      const originalId = task.id;
      task.status = "done";
      task.description = "Updated";

      assert.equal(task.id, originalId);
    });

    it("task createdAt should remain constant", () => {
      const createdAt = nowIso();
      const task: Task = {
        id: 1,
        description: "Test task",
        status: "todo",
        createdAt: createdAt,
        updatedAt: createdAt,
      };

      task.status = "done";
      task.description = "Updated";
      task.updatedAt = nowIso();

      assert.equal(task.createdAt, createdAt);
    });
  });
});
