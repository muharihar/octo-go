// Code generated by octo-go; DO NOT EDIT.

package internal

const (
	AttrRedirectOnly EndpointAttribute = iota
	AttrBoolean
	AttrBodyUploader
	AttrJSONRequestBody
	AttrExplicitURL
	AttrNoResponseBody
)

func init() {
	endpointAttributes = map[string][]EndpointAttribute{

		"ActionsCreateOrUpdateOrgSecretReq":         {AttrJSONRequestBody},
		"ActionsCreateOrUpdateRepoSecretReq":        {AttrJSONRequestBody},
		"ActionsCreateWorkflowDispatchReq":          {AttrJSONRequestBody},
		"ActionsDownloadArtifactReq":                {AttrRedirectOnly},
		"ActionsDownloadJobLogsForWorkflowRunReq":   {AttrRedirectOnly},
		"ActionsDownloadWorkflowRunLogsReq":         {AttrRedirectOnly},
		"ActionsSetSelectedReposForOrgSecretReq":    {AttrJSONRequestBody},
		"ActivityMarkNotificationsAsReadReq":        {AttrJSONRequestBody},
		"ActivityMarkRepoNotificationsAsReadReq":    {AttrJSONRequestBody},
		"ActivitySetRepoSubscriptionReq":            {AttrJSONRequestBody},
		"ActivitySetThreadSubscriptionReq":          {AttrJSONRequestBody},
		"AppsCheckTokenReq":                         {AttrJSONRequestBody},
		"AppsCreateContentAttachmentReq":            {AttrJSONRequestBody},
		"AppsCreateInstallationAccessTokenReq":      {AttrJSONRequestBody},
		"AppsDeleteAuthorizationReq":                {AttrJSONRequestBody},
		"AppsDeleteInstallationReq":                 {AttrBoolean},
		"AppsDeleteTokenReq":                        {AttrJSONRequestBody},
		"AppsResetTokenReq":                         {AttrJSONRequestBody},
		"AppsSuspendInstallationReq":                {AttrBoolean},
		"AppsUnsuspendInstallationReq":              {AttrBoolean},
		"ChecksCreateReq":                           {AttrJSONRequestBody},
		"ChecksCreateSuiteReq":                      {AttrJSONRequestBody},
		"ChecksSetSuitesPreferencesReq":             {AttrJSONRequestBody},
		"ChecksUpdateReq":                           {AttrJSONRequestBody},
		"GistsCreateCommentReq":                     {AttrJSONRequestBody},
		"GistsCreateReq":                            {AttrJSONRequestBody},
		"GistsUpdateCommentReq":                     {AttrJSONRequestBody},
		"GistsUpdateReq":                            {AttrJSONRequestBody},
		"GitCreateBlobReq":                          {AttrJSONRequestBody},
		"GitCreateCommitReq":                        {AttrJSONRequestBody},
		"GitCreateRefReq":                           {AttrJSONRequestBody},
		"GitCreateTagReq":                           {AttrJSONRequestBody},
		"GitCreateTreeReq":                          {AttrJSONRequestBody},
		"GitUpdateRefReq":                           {AttrJSONRequestBody},
		"InteractionsSetRestrictionsForOrgReq":      {AttrJSONRequestBody},
		"InteractionsSetRestrictionsForRepoReq":     {AttrJSONRequestBody},
		"IssuesAddAssigneesReq":                     {AttrJSONRequestBody},
		"IssuesAddLabelsReq":                        {AttrJSONRequestBody},
		"IssuesCheckUserCanBeAssignedReq":           {AttrBoolean},
		"IssuesCreateCommentReq":                    {AttrJSONRequestBody},
		"IssuesCreateLabelReq":                      {AttrJSONRequestBody},
		"IssuesCreateMilestoneReq":                  {AttrJSONRequestBody},
		"IssuesCreateReq":                           {AttrJSONRequestBody},
		"IssuesDeleteMilestoneReq":                  {AttrBoolean},
		"IssuesLockReq":                             {AttrJSONRequestBody},
		"IssuesRemoveAssigneesReq":                  {AttrJSONRequestBody},
		"IssuesSetLabelsReq":                        {AttrJSONRequestBody},
		"IssuesUpdateCommentReq":                    {AttrJSONRequestBody},
		"IssuesUpdateLabelReq":                      {AttrJSONRequestBody},
		"IssuesUpdateMilestoneReq":                  {AttrJSONRequestBody},
		"IssuesUpdateReq":                           {AttrJSONRequestBody},
		"MarkdownRenderRawReq":                      {AttrBodyUploader},
		"MarkdownRenderReq":                         {AttrJSONRequestBody},
		"MigrationsDeleteArchiveForOrgReq":          {AttrBoolean},
		"MigrationsMapCommitAuthorReq":              {AttrJSONRequestBody},
		"MigrationsSetLfsPreferenceReq":             {AttrJSONRequestBody},
		"MigrationsStartForAuthenticatedUserReq":    {AttrJSONRequestBody},
		"MigrationsStartForOrgReq":                  {AttrJSONRequestBody},
		"MigrationsStartImportReq":                  {AttrJSONRequestBody},
		"MigrationsUnlockRepoForOrgReq":             {AttrBoolean},
		"MigrationsUpdateImportReq":                 {AttrJSONRequestBody},
		"OauthAuthorizationsCreateAuthorizationReq": {AttrJSONRequestBody},
		"OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq": {AttrJSONRequestBody},
		"OauthAuthorizationsGetOrCreateAuthorizationForAppReq":               {AttrJSONRequestBody},
		"OauthAuthorizationsUpdateAuthorizationReq":                          {AttrJSONRequestBody},
		"OrgsCheckBlockedUserReq":                                            {AttrBoolean},
		"OrgsCheckPublicMembershipForUserReq":                                {AttrBoolean},
		"OrgsCreateInvitationReq":                                            {AttrJSONRequestBody},
		"OrgsCreateWebhookReq":                                               {AttrJSONRequestBody},
		"OrgsDeleteWebhookReq":                                               {AttrBoolean},
		"OrgsPingWebhookReq":                                                 {AttrBoolean},
		"OrgsRemoveSamlSsoAuthorizationReq":                                  {AttrBoolean},
		"OrgsSetMembershipForUserReq":                                        {AttrJSONRequestBody},
		"OrgsUpdateMembershipForAuthenticatedUserReq":                        {AttrJSONRequestBody},
		"OrgsUpdateReq":                                                      {AttrJSONRequestBody},
		"OrgsUpdateWebhookReq":                                               {AttrJSONRequestBody},
		"ProjectsAddCollaboratorReq":                                         {AttrJSONRequestBody},
		"ProjectsCreateCardReq":                                              {AttrJSONRequestBody},
		"ProjectsCreateColumnReq":                                            {AttrJSONRequestBody},
		"ProjectsCreateForAuthenticatedUserReq":                              {AttrJSONRequestBody},
		"ProjectsCreateForOrgReq":                                            {AttrJSONRequestBody},
		"ProjectsCreateForRepoReq":                                           {AttrJSONRequestBody},
		"ProjectsMoveCardReq":                                                {AttrJSONRequestBody, AttrNoResponseBody},
		"ProjectsMoveColumnReq":                                              {AttrJSONRequestBody, AttrNoResponseBody},
		"ProjectsUpdateCardReq":                                              {AttrJSONRequestBody},
		"ProjectsUpdateColumnReq":                                            {AttrJSONRequestBody},
		"ProjectsUpdateReq":                                                  {AttrJSONRequestBody},
		"PullsCheckIfMergedReq":                                              {AttrBoolean},
		"PullsCreateReplyForReviewCommentReq":                                {AttrJSONRequestBody},
		"PullsCreateReq":                                                     {AttrJSONRequestBody},
		"PullsCreateReviewCommentReq":                                        {AttrJSONRequestBody},
		"PullsCreateReviewReq":                                               {AttrJSONRequestBody},
		"PullsDeleteReviewCommentReq":                                        {AttrBoolean},
		"PullsDismissReviewReq":                                              {AttrJSONRequestBody},
		"PullsMergeReq":                                                      {AttrJSONRequestBody},
		"PullsRemoveRequestedReviewersReq":                                   {AttrJSONRequestBody},
		"PullsRequestReviewersReq":                                           {AttrJSONRequestBody},
		"PullsSubmitReviewReq":                                               {AttrJSONRequestBody},
		"PullsUpdateBranchReq":                                               {AttrJSONRequestBody},
		"PullsUpdateReq":                                                     {AttrJSONRequestBody},
		"PullsUpdateReviewCommentReq":                                        {AttrJSONRequestBody},
		"PullsUpdateReviewReq":                                               {AttrJSONRequestBody},
		"ReactionsCreateForCommitCommentReq":                                 {AttrJSONRequestBody},
		"ReactionsCreateForIssueCommentReq":                                  {AttrJSONRequestBody},
		"ReactionsCreateForIssueReq":                                         {AttrJSONRequestBody},
		"ReactionsCreateForPullRequestReviewCommentReq":                      {AttrJSONRequestBody},
		"ReactionsCreateForTeamDiscussionCommentInOrgReq":                    {AttrJSONRequestBody},
		"ReactionsCreateForTeamDiscussionCommentLegacyReq":                   {AttrJSONRequestBody},
		"ReactionsCreateForTeamDiscussionInOrgReq":                           {AttrJSONRequestBody},
		"ReactionsCreateForTeamDiscussionLegacyReq":                          {AttrJSONRequestBody},
		"ReposAddAppAccessRestrictionsReq":                                   {AttrJSONRequestBody},
		"ReposAddCollaboratorReq":                                            {AttrJSONRequestBody},
		"ReposAddStatusCheckContextsReq":                                     {AttrJSONRequestBody},
		"ReposAddTeamAccessRestrictionsReq":                                  {AttrJSONRequestBody},
		"ReposAddUserAccessRestrictionsReq":                                  {AttrJSONRequestBody},
		"ReposCheckCollaboratorReq":                                          {AttrBoolean},
		"ReposCheckVulnerabilityAlertsReq":                                   {AttrBoolean},
		"ReposCreateCommitCommentReq":                                        {AttrJSONRequestBody},
		"ReposCreateCommitStatusReq":                                         {AttrJSONRequestBody},
		"ReposCreateDeployKeyReq":                                            {AttrJSONRequestBody},
		"ReposCreateDeploymentReq":                                           {AttrJSONRequestBody},
		"ReposCreateDeploymentStatusReq":                                     {AttrJSONRequestBody},
		"ReposCreateDispatchEventReq":                                        {AttrJSONRequestBody},
		"ReposCreateForAuthenticatedUserReq":                                 {AttrJSONRequestBody},
		"ReposCreateForkReq":                                                 {AttrJSONRequestBody},
		"ReposCreateInOrgReq":                                                {AttrJSONRequestBody},
		"ReposCreateOrUpdateFileContentsReq":                                 {AttrJSONRequestBody},
		"ReposCreatePagesSiteReq":                                            {AttrJSONRequestBody},
		"ReposCreateReleaseReq":                                              {AttrJSONRequestBody},
		"ReposCreateUsingTemplateReq":                                        {AttrJSONRequestBody},
		"ReposCreateWebhookReq":                                              {AttrJSONRequestBody},
		"ReposDeleteAdminBranchProtectionReq":                                {AttrBoolean},
		"ReposDeleteCommitCommentReq":                                        {AttrBoolean},
		"ReposDeleteCommitSignatureProtectionReq":                            {AttrBoolean},
		"ReposDeleteFileReq":                                                 {AttrJSONRequestBody},
		"ReposDeletePullRequestReviewProtectionReq":                          {AttrBoolean},
		"ReposDeleteWebhookReq":                                              {AttrBoolean},
		"ReposDownloadTarballArchiveReq":                                     {AttrRedirectOnly},
		"ReposDownloadZipballArchiveReq":                                     {AttrRedirectOnly},
		"ReposMergeReq":                                                      {AttrJSONRequestBody},
		"ReposPingWebhookReq":                                                {AttrBoolean},
		"ReposRemoveAppAccessRestrictionsReq":                                {AttrJSONRequestBody},
		"ReposRemoveStatusCheckContextsReq":                                  {AttrJSONRequestBody},
		"ReposRemoveTeamAccessRestrictionsReq":                               {AttrJSONRequestBody},
		"ReposRemoveUserAccessRestrictionsReq":                               {AttrJSONRequestBody},
		"ReposReplaceAllTopicsReq":                                           {AttrJSONRequestBody},
		"ReposSetAppAccessRestrictionsReq":                                   {AttrJSONRequestBody},
		"ReposSetStatusCheckContextsReq":                                     {AttrJSONRequestBody},
		"ReposSetTeamAccessRestrictionsReq":                                  {AttrJSONRequestBody},
		"ReposSetUserAccessRestrictionsReq":                                  {AttrJSONRequestBody},
		"ReposTestPushWebhookReq":                                            {AttrBoolean},
		"ReposTransferReq":                                                   {AttrJSONRequestBody},
		"ReposUpdateBranchProtectionReq":                                     {AttrJSONRequestBody},
		"ReposUpdateCommitCommentReq":                                        {AttrJSONRequestBody},
		"ReposUpdateInformationAboutPagesSiteReq":                            {AttrJSONRequestBody},
		"ReposUpdateInvitationReq":                                           {AttrJSONRequestBody},
		"ReposUpdatePullRequestReviewProtectionReq":                          {AttrJSONRequestBody},
		"ReposUpdateReleaseAssetReq":                                         {AttrJSONRequestBody},
		"ReposUpdateReleaseReq":                                              {AttrJSONRequestBody},
		"ReposUpdateReq":                                                     {AttrJSONRequestBody},
		"ReposUpdateStatusCheckProtectionReq":                                {AttrJSONRequestBody},
		"ReposUpdateWebhookReq":                                              {AttrJSONRequestBody},
		"ReposUploadReleaseAssetReq":                                         {AttrBodyUploader, AttrExplicitURL},
		"ScimProvisionAndInviteUserReq":                                      {AttrJSONRequestBody},
		"ScimSetInformationForProvisionedUserReq":                            {AttrJSONRequestBody},
		"ScimUpdateAttributeForUserReq":                                      {AttrJSONRequestBody},
		"TeamsAddOrUpdateMembershipForUserInOrgReq":                          {AttrJSONRequestBody},
		"TeamsAddOrUpdateMembershipForUserLegacyReq":                         {AttrJSONRequestBody},
		"TeamsAddOrUpdateProjectPermissionsInOrgReq":                         {AttrJSONRequestBody},
		"TeamsAddOrUpdateProjectPermissionsLegacyReq":                        {AttrJSONRequestBody},
		"TeamsAddOrUpdateRepoPermissionsInOrgReq":                            {AttrJSONRequestBody},
		"TeamsAddOrUpdateRepoPermissionsLegacyReq":                           {AttrJSONRequestBody},
		"TeamsCreateDiscussionCommentInOrgReq":                               {AttrJSONRequestBody},
		"TeamsCreateDiscussionCommentLegacyReq":                              {AttrJSONRequestBody},
		"TeamsCreateDiscussionInOrgReq":                                      {AttrJSONRequestBody},
		"TeamsCreateDiscussionLegacyReq":                                     {AttrJSONRequestBody},
		"TeamsCreateOrUpdateIdpGroupConnectionsInOrgReq":                     {AttrJSONRequestBody},
		"TeamsCreateOrUpdateIdpGroupConnectionsLegacyReq":                    {AttrJSONRequestBody},
		"TeamsCreateReq":                                                     {AttrJSONRequestBody},
		"TeamsGetMemberLegacyReq":                                            {AttrBoolean},
		"TeamsRemoveMemberLegacyReq":                                         {AttrBoolean},
		"TeamsUpdateDiscussionCommentInOrgReq":                               {AttrJSONRequestBody},
		"TeamsUpdateDiscussionCommentLegacyReq":                              {AttrJSONRequestBody},
		"TeamsUpdateDiscussionInOrgReq":                                      {AttrJSONRequestBody},
		"TeamsUpdateDiscussionLegacyReq":                                     {AttrJSONRequestBody},
		"TeamsUpdateInOrgReq":                                                {AttrJSONRequestBody},
		"TeamsUpdateLegacyReq":                                               {AttrJSONRequestBody},
		"UsersAddEmailForAuthenticatedReq":                                   {AttrJSONRequestBody},
		"UsersCheckFollowingForUserReq":                                      {AttrBoolean},
		"UsersCreateGpgKeyForAuthenticatedReq":                               {AttrJSONRequestBody},
		"UsersCreatePublicSshKeyForAuthenticatedReq":                         {AttrJSONRequestBody},
		"UsersDeleteEmailForAuthenticatedReq":                                {AttrJSONRequestBody},
		"UsersSetPrimaryEmailVisibilityForAuthenticatedReq":                  {AttrJSONRequestBody},
		"UsersUpdateAuthenticatedReq":                                        {AttrJSONRequestBody},
	}
}
