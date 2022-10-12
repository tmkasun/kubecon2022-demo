import ballerina/http;

public type CreatedGroup record {|
    *http:Created;
    Group body;
|};

public type CreatedTask record {|
    *http:Created;
    Task body;
|};

public type GroupName record {
    string name?;
};

public type Group record {
    int id?;
    string name?;
};

public type Task record {
    int id?;
    string title?;
    string status?;
    int groupId?;
};

public type InlineResponse2001 record {
    string status?;
};

public type TaskTitle record {
    string title?;
};

public type InlineResponse200 record {
    int groupId?;
};
