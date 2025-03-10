// Code generated by ifacemaker; DO NOT EDIT.

package awsapi

import (
	"context"

	. "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

// CloudWatchLogs provides an interface to the AWS CloudWatchLogs service.
type CloudWatchLogs interface {
	// Associates the specified Key Management Service customer master key (CMK) with
	// the specified log group. Associating an KMS CMK with a log group overrides any
	// existing associations between the log group and a CMK. After a CMK is associated
	// with a log group, all newly ingested data for the log group is encrypted using
	// the CMK. This association is stored as long as the data encrypted with the CMK
	// is still within CloudWatch Logs. This enables CloudWatch Logs to decrypt this
	// data whenever it is requested. CloudWatch Logs supports only symmetric CMKs. Do
	// not use an associate an asymmetric CMK with your log group. For more
	// information, see Using Symmetric and Asymmetric Keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html).
	// It can take up to 5 minutes for this operation to take effect. If you attempt to
	// associate a CMK with a log group but the CMK does not exist or the CMK is
	// disabled, you receive an InvalidParameterException error.
	AssociateKmsKey(ctx context.Context, params *AssociateKmsKeyInput, optFns ...func(*Options)) (*AssociateKmsKeyOutput, error)
	// Cancels the specified export task. The task must be in the PENDING or RUNNING
	// state.
	CancelExportTask(ctx context.Context, params *CancelExportTaskInput, optFns ...func(*Options)) (*CancelExportTaskOutput, error)
	// Creates an export task, which allows you to efficiently export data from a log
	// group to an Amazon S3 bucket. When you perform a CreateExportTask operation, you
	// must use credentials that have permission to write to the S3 bucket that you
	// specify as the destination. Exporting log data to Amazon S3 buckets that are
	// encrypted by KMS is not supported. Exporting log data to Amazon S3 buckets that
	// have S3 Object Lock enabled with a retention period is not supported. Exporting
	// to S3 buckets that are encrypted with AES-256 is supported. This is an
	// asynchronous call. If all the required information is provided, this operation
	// initiates an export task and responds with the ID of the task. After the task
	// has started, you can use DescribeExportTasks
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeExportTasks.html)
	// to get the status of the export task. Each account can only have one active
	// (RUNNING or PENDING) export task at a time. To cancel an export task, use
	// CancelExportTask
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CancelExportTask.html).
	// You can export logs from multiple log groups or multiple time ranges to the same
	// S3 bucket. To separate out log data for each export task, you can specify a
	// prefix to be used as the Amazon S3 key prefix for all exported objects.
	// Time-based sorting on chunks of log data inside an exported file is not
	// guaranteed. You can sort the exported log fild data by using Linux utilities.
	CreateExportTask(ctx context.Context, params *CreateExportTaskInput, optFns ...func(*Options)) (*CreateExportTaskOutput, error)
	// Creates a log group with the specified name. You can create up to 20,000 log
	// groups per account. You must use the following guidelines when naming a log
	// group:
	//
	// * Log group names must be unique within a region for an Amazon Web
	// Services account.
	//
	// * Log group names can be between 1 and 512 characters
	// long.
	//
	// * Log group names consist of the following characters: a-z, A-Z, 0-9, '_'
	// (underscore), '-' (hyphen), '/' (forward slash), '.' (period), and '#' (number
	// sign)
	//
	// When you create a log group, by default the log events in the log group
	// never expire. To set a retention policy so that events expire and are deleted
	// after a specified time, use PutRetentionPolicy
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutRetentionPolicy.html).
	// If you associate a Key Management Service customer master key (CMK) with the log
	// group, ingested data is encrypted using the CMK. This association is stored as
	// long as the data encrypted with the CMK is still within CloudWatch Logs. This
	// enables CloudWatch Logs to decrypt this data whenever it is requested. If you
	// attempt to associate a CMK with the log group but the CMK does not exist or the
	// CMK is disabled, you receive an InvalidParameterException error. CloudWatch Logs
	// supports only symmetric CMKs. Do not associate an asymmetric CMK with your log
	// group. For more information, see Using Symmetric and Asymmetric Keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html).
	CreateLogGroup(ctx context.Context, params *CreateLogGroupInput, optFns ...func(*Options)) (*CreateLogGroupOutput, error)
	// Creates a log stream for the specified log group. A log stream is a sequence of
	// log events that originate from a single source, such as an application instance
	// or a resource that is being monitored. There is no limit on the number of log
	// streams that you can create for a log group. There is a limit of 50 TPS on
	// CreateLogStream operations, after which transactions are throttled. You must use
	// the following guidelines when naming a log stream:
	//
	// * Log stream names must be
	// unique within the log group.
	//
	// * Log stream names can be between 1 and 512
	// characters long.
	//
	// * The ':' (colon) and '*' (asterisk) characters are not
	// allowed.
	CreateLogStream(ctx context.Context, params *CreateLogStreamInput, optFns ...func(*Options)) (*CreateLogStreamOutput, error)
	// Deletes the specified destination, and eventually disables all the subscription
	// filters that publish to it. This operation does not delete the physical resource
	// encapsulated by the destination.
	DeleteDestination(ctx context.Context, params *DeleteDestinationInput, optFns ...func(*Options)) (*DeleteDestinationOutput, error)
	// Deletes the specified log group and permanently deletes all the archived log
	// events associated with the log group.
	DeleteLogGroup(ctx context.Context, params *DeleteLogGroupInput, optFns ...func(*Options)) (*DeleteLogGroupOutput, error)
	// Deletes the specified log stream and permanently deletes all the archived log
	// events associated with the log stream.
	DeleteLogStream(ctx context.Context, params *DeleteLogStreamInput, optFns ...func(*Options)) (*DeleteLogStreamOutput, error)
	// Deletes the specified metric filter.
	DeleteMetricFilter(ctx context.Context, params *DeleteMetricFilterInput, optFns ...func(*Options)) (*DeleteMetricFilterOutput, error)
	// Deletes a saved CloudWatch Logs Insights query definition. A query definition
	// contains details about a saved CloudWatch Logs Insights query. Each
	// DeleteQueryDefinition operation can delete one query definition. You must have
	// the logs:DeleteQueryDefinition permission to be able to perform this operation.
	DeleteQueryDefinition(ctx context.Context, params *DeleteQueryDefinitionInput, optFns ...func(*Options)) (*DeleteQueryDefinitionOutput, error)
	// Deletes a resource policy from this account. This revokes the access of the
	// identities in that policy to put log events to this account.
	DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error)
	// Deletes the specified retention policy. Log events do not expire if they belong
	// to log groups without a retention policy.
	DeleteRetentionPolicy(ctx context.Context, params *DeleteRetentionPolicyInput, optFns ...func(*Options)) (*DeleteRetentionPolicyOutput, error)
	// Deletes the specified subscription filter.
	DeleteSubscriptionFilter(ctx context.Context, params *DeleteSubscriptionFilterInput, optFns ...func(*Options)) (*DeleteSubscriptionFilterOutput, error)
	// Lists all your destinations. The results are ASCII-sorted by destination name.
	DescribeDestinations(ctx context.Context, params *DescribeDestinationsInput, optFns ...func(*Options)) (*DescribeDestinationsOutput, error)
	// Lists the specified export tasks. You can list all your export tasks or filter
	// the results based on task ID or task status.
	DescribeExportTasks(ctx context.Context, params *DescribeExportTasksInput, optFns ...func(*Options)) (*DescribeExportTasksOutput, error)
	// Lists the specified log groups. You can list all your log groups or filter the
	// results by prefix. The results are ASCII-sorted by log group name. CloudWatch
	// Logs doesn’t support IAM policies that control access to the DescribeLogGroups
	// action by using the aws:ResourceTag/key-name  condition key. Other CloudWatch
	// Logs actions do support the use of the aws:ResourceTag/key-name  condition key
	// to control access. For more information about using tags to control access, see
	// Controlling access to Amazon Web Services resources using tags
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html).
	DescribeLogGroups(ctx context.Context, params *DescribeLogGroupsInput, optFns ...func(*Options)) (*DescribeLogGroupsOutput, error)
	// Lists the log streams for the specified log group. You can list all the log
	// streams or filter the results by prefix. You can also control how the results
	// are ordered. This operation has a limit of five transactions per second, after
	// which transactions are throttled.
	DescribeLogStreams(ctx context.Context, params *DescribeLogStreamsInput, optFns ...func(*Options)) (*DescribeLogStreamsOutput, error)
	// Lists the specified metric filters. You can list all of the metric filters or
	// filter the results by log name, prefix, metric name, or metric namespace. The
	// results are ASCII-sorted by filter name.
	DescribeMetricFilters(ctx context.Context, params *DescribeMetricFiltersInput, optFns ...func(*Options)) (*DescribeMetricFiltersOutput, error)
	// Returns a list of CloudWatch Logs Insights queries that are scheduled,
	// executing, or have been executed recently in this account. You can request all
	// queries or limit it to queries of a specific log group or queries with a certain
	// status.
	DescribeQueries(ctx context.Context, params *DescribeQueriesInput, optFns ...func(*Options)) (*DescribeQueriesOutput, error)
	// This operation returns a paginated list of your saved CloudWatch Logs Insights
	// query definitions. You can use the queryDefinitionNamePrefix parameter to limit
	// the results to only the query definitions that have names that start with a
	// certain string.
	DescribeQueryDefinitions(ctx context.Context, params *DescribeQueryDefinitionsInput, optFns ...func(*Options)) (*DescribeQueryDefinitionsOutput, error)
	// Lists the resource policies in this account.
	DescribeResourcePolicies(ctx context.Context, params *DescribeResourcePoliciesInput, optFns ...func(*Options)) (*DescribeResourcePoliciesOutput, error)
	// Lists the subscription filters for the specified log group. You can list all the
	// subscription filters or filter the results by prefix. The results are
	// ASCII-sorted by filter name.
	DescribeSubscriptionFilters(ctx context.Context, params *DescribeSubscriptionFiltersInput, optFns ...func(*Options)) (*DescribeSubscriptionFiltersOutput, error)
	// Disassociates the associated Key Management Service customer master key (CMK)
	// from the specified log group. After the KMS CMK is disassociated from the log
	// group, CloudWatch Logs stops encrypting newly ingested data for the log group.
	// All previously ingested data remains encrypted, and CloudWatch Logs requires
	// permissions for the CMK whenever the encrypted data is requested. Note that it
	// can take up to 5 minutes for this operation to take effect.
	DisassociateKmsKey(ctx context.Context, params *DisassociateKmsKeyInput, optFns ...func(*Options)) (*DisassociateKmsKeyOutput, error)
	// Lists log events from the specified log group. You can list all the log events
	// or filter the results using a filter pattern, a time range, and the name of the
	// log stream. By default, this operation returns as many log events as can fit in
	// 1 MB (up to 10,000 log events) or all the events found within the time range
	// that you specify. If the results include a token, then there are more log events
	// available, and you can get additional results by specifying the token in a
	// subsequent call. This operation can return empty results while there are more
	// log events available through the token. The returned log events are sorted by
	// event timestamp, the timestamp when the event was ingested by CloudWatch Logs,
	// and the ID of the PutLogEvents request.
	FilterLogEvents(ctx context.Context, params *FilterLogEventsInput, optFns ...func(*Options)) (*FilterLogEventsOutput, error)
	// Lists log events from the specified log stream. You can list all of the log
	// events or filter using a time range. By default, this operation returns as many
	// log events as can fit in a response size of 1MB (up to 10,000 log events). You
	// can get additional log events by specifying one of the tokens in a subsequent
	// call. This operation can return empty results while there are more log events
	// available through the token.
	GetLogEvents(ctx context.Context, params *GetLogEventsInput, optFns ...func(*Options)) (*GetLogEventsOutput, error)
	// Returns a list of the fields that are included in log events in the specified
	// log group, along with the percentage of log events that contain each field. The
	// search is limited to a time period that you specify. In the results, fields that
	// start with @ are fields generated by CloudWatch Logs. For example, @timestamp is
	// the timestamp of each log event. For more information about the fields that are
	// generated by CloudWatch logs, see Supported Logs and Discovered Fields
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_AnalyzeLogData-discoverable-fields.html).
	// The response results are sorted by the frequency percentage, starting with the
	// highest percentage.
	GetLogGroupFields(ctx context.Context, params *GetLogGroupFieldsInput, optFns ...func(*Options)) (*GetLogGroupFieldsOutput, error)
	// Retrieves all of the fields and values of a single log event. All fields are
	// retrieved, even if the original query that produced the logRecordPointer
	// retrieved only a subset of fields. Fields are returned as field name/field value
	// pairs. The full unparsed log event is returned within @message.
	GetLogRecord(ctx context.Context, params *GetLogRecordInput, optFns ...func(*Options)) (*GetLogRecordOutput, error)
	// Returns the results from the specified query. Only the fields requested in the
	// query are returned, along with a @ptr field, which is the identifier for the log
	// record. You can use the value of @ptr in a GetLogRecord
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogRecord.html)
	// operation to get the full log record. GetQueryResults does not start a query
	// execution. To run a query, use StartQuery
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartQuery.html).
	// If the value of the Status field in the output is Running, this operation
	// returns only partial results. If you see a value of Scheduled or Running for the
	// status, you can retry the operation later to see the final results.
	GetQueryResults(ctx context.Context, params *GetQueryResultsInput, optFns ...func(*Options)) (*GetQueryResultsOutput, error)
	// Lists the tags for the specified log group.
	ListTagsLogGroup(ctx context.Context, params *ListTagsLogGroupInput, optFns ...func(*Options)) (*ListTagsLogGroupOutput, error)
	// Creates or updates a destination. This operation is used only to create
	// destinations for cross-account subscriptions. A destination encapsulates a
	// physical resource (such as an Amazon Kinesis stream) and enables you to
	// subscribe to a real-time stream of log events for a different account, ingested
	// using PutLogEvents
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html).
	// Through an access policy, a destination controls what is written to it. By
	// default, PutDestination does not set any access policy with the destination,
	// which means a cross-account user cannot call PutSubscriptionFilter
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutSubscriptionFilter.html)
	// against this destination. To enable this, the destination owner must call
	// PutDestinationPolicy
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDestinationPolicy.html)
	// after PutDestination. To perform a PutDestination operation, you must also have
	// the iam:PassRole permission.
	PutDestination(ctx context.Context, params *PutDestinationInput, optFns ...func(*Options)) (*PutDestinationOutput, error)
	// Creates or updates an access policy associated with an existing destination. An
	// access policy is an IAM policy document
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies_overview.html) that
	// is used to authorize claims to register a subscription filter against a given
	// destination. If multiple Amazon Web Services accounts are sending logs to this
	// destination, each sender account must be listed separately in the policy. The
	// policy does not support specifying * as the Principal or the use of the
	// aws:PrincipalOrgId global key.
	PutDestinationPolicy(ctx context.Context, params *PutDestinationPolicyInput, optFns ...func(*Options)) (*PutDestinationPolicyOutput, error)
	// Uploads a batch of log events to the specified log stream. You must include the
	// sequence token obtained from the response of the previous call. An upload in a
	// newly created log stream does not require a sequence token. You can also get the
	// sequence token in the expectedSequenceToken field from
	// InvalidSequenceTokenException. If you call PutLogEvents twice within a narrow
	// time period using the same value for sequenceToken, both calls might be
	// successful or one might be rejected. The batch of events must satisfy the
	// following constraints:
	//
	// * The maximum batch size is 1,048,576 bytes. This size
	// is calculated as the sum of all event messages in UTF-8, plus 26 bytes for each
	// log event.
	//
	// * None of the log events in the batch can be more than 2 hours in
	// the future.
	//
	// * None of the log events in the batch can be older than 14 days or
	// older than the retention period of the log group.
	//
	// * The log events in the batch
	// must be in chronological order by their timestamp. The timestamp is the time the
	// event occurred, expressed as the number of milliseconds after Jan 1, 1970
	// 00:00:00 UTC. (In Amazon Web Services Tools for PowerShell and the Amazon Web
	// Services SDK for .NET, the timestamp is specified in .NET format:
	// yyyy-mm-ddThh:mm:ss. For example, 2017-09-15T13:45:30.)
	//
	// * A batch of log events
	// in a single request cannot span more than 24 hours. Otherwise, the operation
	// fails.
	//
	// * The maximum number of log events in a batch is 10,000.
	//
	// * There is a
	// quota of 5 requests per second per log stream. Additional requests are
	// throttled. This quota can't be changed.
	//
	// If a call to PutLogEvents returns
	// "UnrecognizedClientException" the most likely cause is an invalid Amazon Web
	// Services access key ID or secret key.
	PutLogEvents(ctx context.Context, params *PutLogEventsInput, optFns ...func(*Options)) (*PutLogEventsOutput, error)
	// Creates or updates a metric filter and associates it with the specified log
	// group. Metric filters allow you to configure rules to extract metric data from
	// log events ingested through PutLogEvents
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html).
	// The maximum number of metric filters that can be associated with a log group is
	// 100. When you create a metric filter, you can also optionally assign a unit and
	// dimensions to the metric that is created. Metrics extracted from log events are
	// charged as custom metrics. To prevent unexpected high charges, do not specify
	// high-cardinality fields such as IPAddress or requestID as dimensions. Each
	// different value found for a dimension is treated as a separate metric and
	// accrues charges as a separate custom metric. To help prevent accidental high
	// charges, Amazon disables a metric filter if it generates 1000 different
	// name/value pairs for the dimensions that you have specified within a certain
	// amount of time. You can also set up a billing alarm to alert you if your charges
	// are higher than expected. For more information, see  Creating a Billing Alarm to
	// Monitor Your Estimated Amazon Web Services Charges
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/monitor_estimated_charges_with_cloudwatch.html).
	PutMetricFilter(ctx context.Context, params *PutMetricFilterInput, optFns ...func(*Options)) (*PutMetricFilterOutput, error)
	// Creates or updates a query definition for CloudWatch Logs Insights. For more
	// information, see Analyzing Log Data with CloudWatch Logs Insights
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AnalyzingLogData.html).
	// To update a query definition, specify its queryDefinitionId in your request. The
	// values of name, queryString, and logGroupNames are changed to the values that
	// you specify in your update operation. No current values are retained from the
	// current query definition. For example, if you update a current query definition
	// that includes log groups, and you don't specify the logGroupNames parameter in
	// your update operation, the query definition changes to contain no log groups.
	// You must have the logs:PutQueryDefinition permission to be able to perform this
	// operation.
	PutQueryDefinition(ctx context.Context, params *PutQueryDefinitionInput, optFns ...func(*Options)) (*PutQueryDefinitionOutput, error)
	// Creates or updates a resource policy allowing other Amazon Web Services services
	// to put log events to this account, such as Amazon Route 53. An account can have
	// up to 10 resource policies per Amazon Web Services Region.
	PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error)
	// Sets the retention of the specified log group. A retention policy allows you to
	// configure the number of days for which to retain log events in the specified log
	// group.
	PutRetentionPolicy(ctx context.Context, params *PutRetentionPolicyInput, optFns ...func(*Options)) (*PutRetentionPolicyOutput, error)
	// Creates or updates a subscription filter and associates it with the specified
	// log group. Subscription filters allow you to subscribe to a real-time stream of
	// log events ingested through PutLogEvents
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html)
	// and have them delivered to a specific destination. When log events are sent to
	// the receiving service, they are Base64 encoded and compressed with the gzip
	// format. The following destinations are supported for subscription filters:
	//
	// * An
	// Amazon Kinesis stream belonging to the same account as the subscription filter,
	// for same-account delivery.
	//
	// * A logical destination that belongs to a different
	// account, for cross-account delivery.
	//
	// * An Amazon Kinesis Firehose delivery
	// stream that belongs to the same account as the subscription filter, for
	// same-account delivery.
	//
	// * An Lambda function that belongs to the same account as
	// the subscription filter, for same-account delivery.
	//
	// Each log group can have up
	// to two subscription filters associated with it. If you are updating an existing
	// filter, you must specify the correct name in filterName. To perform a
	// PutSubscriptionFilter operation, you must also have the iam:PassRole permission.
	PutSubscriptionFilter(ctx context.Context, params *PutSubscriptionFilterInput, optFns ...func(*Options)) (*PutSubscriptionFilterOutput, error)
	// Schedules a query of a log group using CloudWatch Logs Insights. You specify the
	// log group and time range to query and the query string to use. For more
	// information, see CloudWatch Logs Insights Query Syntax
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
	// Queries time out after 15 minutes of execution. If your queries are timing out,
	// reduce the time range being searched or partition your query into a number of
	// queries.
	StartQuery(ctx context.Context, params *StartQueryInput, optFns ...func(*Options)) (*StartQueryOutput, error)
	// Stops a CloudWatch Logs Insights query that is in progress. If the query has
	// already ended, the operation returns an error indicating that the specified
	// query is not running.
	StopQuery(ctx context.Context, params *StopQueryInput, optFns ...func(*Options)) (*StopQueryOutput, error)
	// Adds or updates the specified tags for the specified log group. To list the tags
	// for a log group, use ListTagsLogGroup
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ListTagsLogGroup.html).
	// To remove tags, use UntagLogGroup
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_UntagLogGroup.html).
	// For more information about tags, see Tag Log Groups in Amazon CloudWatch Logs
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html#log-group-tagging)
	// in the Amazon CloudWatch Logs User Guide. CloudWatch Logs doesn’t support IAM
	// policies that prevent users from assigning specified tags to log groups using
	// the aws:Resource/key-name  or aws:TagKeys condition keys. For more information
	// about using tags to control access, see Controlling access to Amazon Web
	// Services resources using tags
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html).
	TagLogGroup(ctx context.Context, params *TagLogGroupInput, optFns ...func(*Options)) (*TagLogGroupOutput, error)
	// Tests the filter pattern of a metric filter against a sample of log event
	// messages. You can use this operation to validate the correctness of a metric
	// filter pattern.
	TestMetricFilter(ctx context.Context, params *TestMetricFilterInput, optFns ...func(*Options)) (*TestMetricFilterOutput, error)
	// Removes the specified tags from the specified log group. To list the tags for a
	// log group, use ListTagsLogGroup
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ListTagsLogGroup.html).
	// To add tags, use TagLogGroup
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_TagLogGroup.html).
	// CloudWatch Logs doesn’t support IAM policies that prevent users from assigning
	// specified tags to log groups using the aws:Resource/key-name  or aws:TagKeys
	// condition keys.
	UntagLogGroup(ctx context.Context, params *UntagLogGroupInput, optFns ...func(*Options)) (*UntagLogGroupOutput, error)
}

