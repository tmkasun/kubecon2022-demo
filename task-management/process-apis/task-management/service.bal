import ballerina/http;
import task_management.entity;

listener http:Listener ep0 = new (9090, config = {host: "localhost"});

service /task/manage on ep0 {
    resource function post tasks(@http:Payload entity:TasksBody payload) returns entity:CreatedTask {
      return entity:createNewTask(payload);
    }
    resource function put tasks/[string taskId](@http:Payload entity:TasksTaskidBody payload) returns entity:Task|error {
      return entity:updateTask(taskId, payload);
    }
    resource function delete tasks/[string taskId]() returns http:Ok {
      return entity:deleteTask(taskId);
    }
    resource function post tasks/[string taskId]/'change\-group(@http:Payload entity:TaskidChangegroupBody payload) returns entity:InlineResponse200 {
      return entity:changeTaskGroup(taskId, payload);
    }
    resource function post tasks/[string taskId]/'change\-status(@http:Payload entity:TaskidChangestatusBody payload) returns entity:InlineResponse2001 {
      return entity:changeTaskStatus(taskId, payload);
    }
    resource function post groups(@http:Payload entity:GroupName payload) returns entity:CreatedGroup {
      return entity:createNewGroup(payload);
    }
    resource function put groups/[string groupId](@http:Payload entity:GroupName payload) returns entity:Group|error {
      return entity:updateGroup(groupId, payload);
    }
    resource function post archive/tasks(@http:Payload entity:ArchiveTasksBody payload) returns http:Ok {
      return entity:archiveTask(payload);
    }
    resource function post archive/groups(@http:Payload entity:ArchiveGroupsBody payload) returns http:Ok {
      return entity:archiveGroup(payload);
    }
}
