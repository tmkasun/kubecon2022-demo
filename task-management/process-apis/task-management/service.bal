import ballerina/http;

listener http:Listener ep0 = new (9090, config = {host: "localhost"});

service /task/manage on ep0 {
    resource function post tasks(int groupId, @http:Payload TaskTitle payload) returns CreatedTask {
    }
    resource function put tasks/[string taskId](int groupId, @http:Payload TaskTitle payload) returns Task {
    }
    resource function delete tasks/[string taskId](int groupId) returns http:Ok {
    }
    resource function post tasks/[string taskId]/'change\-group(int groupId, int newGroupId) returns InlineResponse200 {
    }
    resource function post tasks/[string taskId]/'change\-status(int groupId, string status) returns InlineResponse2001 {
    }
    resource function post groups(@http:Payload GroupName payload) returns CreatedGroup {
    }
    resource function put groups/[string groupId](@http:Payload GroupName payload) returns Group {
    }
    resource function post archive/tasks(int groupId, int taskId) returns http:Ok {
    }
    resource function post archive/groups(int groupId) returns http:Ok {
    }
}
