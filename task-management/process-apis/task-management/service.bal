import ballerina/http;
import task_management.entity;

listener http:Listener ep0 = new (9090, config = {host: "localhost"});

service /task/manage on ep0 {
    resource function post tasks(int groupId, @http:Payload entity:TaskTitle payload) returns entity:CreatedTask {
      return entity:createNewTask(groupId, payload);
    }
    resource function put tasks/[string taskId](int groupId, @http:Payload entity:TaskTitle payload) returns entity:Task|error {
      return check entity:updateTask(taskId, groupId, payload);
    }
    resource function delete tasks/[string taskId](int groupId) returns http:Ok {
      return entity:deleteTask(taskId);
    }
    resource function post tasks/[string taskId]/'change\-group(int groupId, int newGroupId) returns entity:InlineResponse200 {
      return entity:changeTaskGroup(taskId, groupId, newGroupId);
    }
    resource function post tasks/[string taskId]/'change\-status(int groupId, string status) returns entity:InlineResponse2001 {
      return entity:changeTaskStatus(groupId, status);
    }
    resource function post groups(@http:Payload entity:GroupName payload) returns entity:CreatedGroup {
      return entity:createNewGroup(payload);
    }
    resource function put groups/[string groupId](@http:Payload entity:GroupName payload) returns entity:Group|error {
      return entity:updateGroup(groupId, payload);
    }
    resource function post archive/tasks(int groupId, int taskId) returns http:Ok {
      return entity:archiveTask(groupId, taskId);
    }
    resource function post archive/groups(int groupId) returns http:Ok {
      return entity:archiveGroup(groupId);
    }
}
