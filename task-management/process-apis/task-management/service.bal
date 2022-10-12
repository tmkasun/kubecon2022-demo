import ballerina/http;
import task_management.entity;

listener http:Listener ep0 = new (9090, config = {host: "localhost"});

service /task/manage on ep0 {
    resource function post tasks(int groupId, @http:Payload entity:TaskTitle payload) returns entity:CreatedTask {
    }
    resource function put tasks/[string taskId](int groupId, @http:Payload entity:TaskTitle payload) returns entity:Task {
    }
    resource function delete tasks/[string taskId](int groupId) returns http:Ok {
    }
    resource function post tasks/[string taskId]/'change\-group(int groupId, int newGroupId) returns entity:InlineResponse200 {
    }
    resource function post tasks/[string taskId]/'change\-status(int groupId, string status) returns entity:InlineResponse2001 {
    }
    resource function post groups(@http:Payload entity:GroupName payload) returns entity:CreatedGroup {
    }
    resource function put groups/[string groupId](@http:Payload entity:GroupName payload) returns entity:Group {
    }
    resource function post archive/tasks(int groupId, int taskId) returns http:Ok {
    }
    resource function post archive/groups(int groupId) returns http:Ok {
    }
}
