// Code generated by smithy-go-codegen DO NOT EDIT.

package amplifybackend

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/amplifybackend/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCloneBackend struct {
}

func (*validateOpCloneBackend) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCloneBackend) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CloneBackendInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCloneBackendInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateBackendAPI struct {
}

func (*validateOpCreateBackendAPI) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateBackendAPI) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateBackendAPIInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateBackendAPIInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateBackendAuth struct {
}

func (*validateOpCreateBackendAuth) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateBackendAuth) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateBackendAuthInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateBackendAuthInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateBackendConfig struct {
}

func (*validateOpCreateBackendConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateBackendConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateBackendConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateBackendConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateBackend struct {
}

func (*validateOpCreateBackend) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateBackend) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateBackendInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateBackendInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateToken struct {
}

func (*validateOpCreateToken) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateToken) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateTokenInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateTokenInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteBackendAPI struct {
}

func (*validateOpDeleteBackendAPI) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteBackendAPI) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteBackendAPIInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteBackendAPIInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteBackendAuth struct {
}

func (*validateOpDeleteBackendAuth) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteBackendAuth) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteBackendAuthInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteBackendAuthInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteBackend struct {
}

func (*validateOpDeleteBackend) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteBackend) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteBackendInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteBackendInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteToken struct {
}

func (*validateOpDeleteToken) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteToken) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteTokenInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteTokenInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGenerateBackendAPIModels struct {
}

func (*validateOpGenerateBackendAPIModels) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGenerateBackendAPIModels) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GenerateBackendAPIModelsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGenerateBackendAPIModelsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetBackendAPI struct {
}

func (*validateOpGetBackendAPI) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetBackendAPI) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetBackendAPIInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetBackendAPIInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetBackendAPIModels struct {
}

func (*validateOpGetBackendAPIModels) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetBackendAPIModels) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetBackendAPIModelsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetBackendAPIModelsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetBackendAuth struct {
}

func (*validateOpGetBackendAuth) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetBackendAuth) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetBackendAuthInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetBackendAuthInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetBackend struct {
}

func (*validateOpGetBackend) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetBackend) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetBackendInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetBackendInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetBackendJob struct {
}

func (*validateOpGetBackendJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetBackendJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetBackendJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetBackendJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetToken struct {
}

func (*validateOpGetToken) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetToken) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTokenInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTokenInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListBackendJobs struct {
}

func (*validateOpListBackendJobs) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListBackendJobs) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListBackendJobsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListBackendJobsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRemoveAllBackends struct {
}

func (*validateOpRemoveAllBackends) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRemoveAllBackends) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RemoveAllBackendsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRemoveAllBackendsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRemoveBackendConfig struct {
}

func (*validateOpRemoveBackendConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRemoveBackendConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RemoveBackendConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRemoveBackendConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateBackendAPI struct {
}

func (*validateOpUpdateBackendAPI) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateBackendAPI) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateBackendAPIInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateBackendAPIInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateBackendAuth struct {
}

func (*validateOpUpdateBackendAuth) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateBackendAuth) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateBackendAuthInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateBackendAuthInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateBackendConfig struct {
}

func (*validateOpUpdateBackendConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateBackendConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateBackendConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateBackendConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateBackendJob struct {
}

func (*validateOpUpdateBackendJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateBackendJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateBackendJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateBackendJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCloneBackendValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCloneBackend{}, middleware.After)
}

func addOpCreateBackendAPIValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateBackendAPI{}, middleware.After)
}

func addOpCreateBackendAuthValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateBackendAuth{}, middleware.After)
}

func addOpCreateBackendConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateBackendConfig{}, middleware.After)
}

func addOpCreateBackendValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateBackend{}, middleware.After)
}

func addOpCreateTokenValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateToken{}, middleware.After)
}

func addOpDeleteBackendAPIValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteBackendAPI{}, middleware.After)
}

func addOpDeleteBackendAuthValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteBackendAuth{}, middleware.After)
}

func addOpDeleteBackendValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteBackend{}, middleware.After)
}

func addOpDeleteTokenValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteToken{}, middleware.After)
}

func addOpGenerateBackendAPIModelsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGenerateBackendAPIModels{}, middleware.After)
}

func addOpGetBackendAPIValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetBackendAPI{}, middleware.After)
}

func addOpGetBackendAPIModelsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetBackendAPIModels{}, middleware.After)
}

func addOpGetBackendAuthValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetBackendAuth{}, middleware.After)
}

func addOpGetBackendValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetBackend{}, middleware.After)
}

func addOpGetBackendJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetBackendJob{}, middleware.After)
}

func addOpGetTokenValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetToken{}, middleware.After)
}

func addOpListBackendJobsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListBackendJobs{}, middleware.After)
}

func addOpRemoveAllBackendsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemoveAllBackends{}, middleware.After)
}

func addOpRemoveBackendConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemoveBackendConfig{}, middleware.After)
}

func addOpUpdateBackendAPIValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateBackendAPI{}, middleware.After)
}

func addOpUpdateBackendAuthValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateBackendAuth{}, middleware.After)
}

func addOpUpdateBackendConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateBackendConfig{}, middleware.After)
}

func addOpUpdateBackendJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateBackendJob{}, middleware.After)
}

func validateCreateBackendAuthForgotPasswordConfig(v *types.CreateBackendAuthForgotPasswordConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateBackendAuthForgotPasswordConfig"}
	if len(v.DeliveryMethod) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("DeliveryMethod"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateBackendAuthIdentityPoolConfig(v *types.CreateBackendAuthIdentityPoolConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateBackendAuthIdentityPoolConfig"}
	if v.IdentityPoolName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateBackendAuthMFAConfig(v *types.CreateBackendAuthMFAConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateBackendAuthMFAConfig"}
	if len(v.MFAMode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("MFAMode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateBackendAuthOAuthConfig(v *types.CreateBackendAuthOAuthConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateBackendAuthOAuthConfig"}
	if len(v.OAuthGrantType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("OAuthGrantType"))
	}
	if v.OAuthScopes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OAuthScopes"))
	}
	if v.RedirectSignInURIs == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RedirectSignInURIs"))
	}
	if v.RedirectSignOutURIs == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RedirectSignOutURIs"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateBackendAuthPasswordPolicyConfig(v *types.CreateBackendAuthPasswordPolicyConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateBackendAuthPasswordPolicyConfig"}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateBackendAuthResourceConfig(v *types.CreateBackendAuthResourceConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateBackendAuthResourceConfig"}
	if len(v.AuthResources) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("AuthResources"))
	}
	if v.IdentityPoolConfigs != nil {
		if err := validateCreateBackendAuthIdentityPoolConfig(v.IdentityPoolConfigs); err != nil {
			invalidParams.AddNested("IdentityPoolConfigs", err.(smithy.InvalidParamsError))
		}
	}
	if len(v.Service) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Service"))
	}
	if v.UserPoolConfigs == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserPoolConfigs"))
	} else if v.UserPoolConfigs != nil {
		if err := validateCreateBackendAuthUserPoolConfig(v.UserPoolConfigs); err != nil {
			invalidParams.AddNested("UserPoolConfigs", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateBackendAuthUserPoolConfig(v *types.CreateBackendAuthUserPoolConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateBackendAuthUserPoolConfig"}
	if v.ForgotPassword != nil {
		if err := validateCreateBackendAuthForgotPasswordConfig(v.ForgotPassword); err != nil {
			invalidParams.AddNested("ForgotPassword", err.(smithy.InvalidParamsError))
		}
	}
	if v.Mfa != nil {
		if err := validateCreateBackendAuthMFAConfig(v.Mfa); err != nil {
			invalidParams.AddNested("Mfa", err.(smithy.InvalidParamsError))
		}
	}
	if v.OAuth != nil {
		if err := validateCreateBackendAuthOAuthConfig(v.OAuth); err != nil {
			invalidParams.AddNested("OAuth", err.(smithy.InvalidParamsError))
		}
	}
	if v.PasswordPolicy != nil {
		if err := validateCreateBackendAuthPasswordPolicyConfig(v.PasswordPolicy); err != nil {
			invalidParams.AddNested("PasswordPolicy", err.(smithy.InvalidParamsError))
		}
	}
	if v.RequiredSignUpAttributes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RequiredSignUpAttributes"))
	}
	if len(v.SignInMethod) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("SignInMethod"))
	}
	if v.UserPoolName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserPoolName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateBackendAuthResourceConfig(v *types.UpdateBackendAuthResourceConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateBackendAuthResourceConfig"}
	if len(v.AuthResources) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("AuthResources"))
	}
	if len(v.Service) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Service"))
	}
	if v.UserPoolConfigs == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserPoolConfigs"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCloneBackendInput(v *CloneBackendInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CloneBackendInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.BackendEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackendEnvironmentName"))
	}
	if v.TargetEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetEnvironmentName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateBackendAPIInput(v *CreateBackendAPIInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateBackendAPIInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.BackendEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackendEnvironmentName"))
	}
	if v.ResourceConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceConfig"))
	}
	if v.ResourceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateBackendAuthInput(v *CreateBackendAuthInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateBackendAuthInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.BackendEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackendEnvironmentName"))
	}
	if v.ResourceConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceConfig"))
	} else if v.ResourceConfig != nil {
		if err := validateCreateBackendAuthResourceConfig(v.ResourceConfig); err != nil {
			invalidParams.AddNested("ResourceConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.ResourceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateBackendConfigInput(v *CreateBackendConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateBackendConfigInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateBackendInput(v *CreateBackendInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateBackendInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.AppName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppName"))
	}
	if v.BackendEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackendEnvironmentName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateTokenInput(v *CreateTokenInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateTokenInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteBackendAPIInput(v *DeleteBackendAPIInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteBackendAPIInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.BackendEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackendEnvironmentName"))
	}
	if v.ResourceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteBackendAuthInput(v *DeleteBackendAuthInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteBackendAuthInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.BackendEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackendEnvironmentName"))
	}
	if v.ResourceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteBackendInput(v *DeleteBackendInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteBackendInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.BackendEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackendEnvironmentName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteTokenInput(v *DeleteTokenInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteTokenInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGenerateBackendAPIModelsInput(v *GenerateBackendAPIModelsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GenerateBackendAPIModelsInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.BackendEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackendEnvironmentName"))
	}
	if v.ResourceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetBackendAPIInput(v *GetBackendAPIInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetBackendAPIInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.BackendEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackendEnvironmentName"))
	}
	if v.ResourceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetBackendAPIModelsInput(v *GetBackendAPIModelsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetBackendAPIModelsInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.BackendEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackendEnvironmentName"))
	}
	if v.ResourceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetBackendAuthInput(v *GetBackendAuthInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetBackendAuthInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.BackendEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackendEnvironmentName"))
	}
	if v.ResourceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetBackendInput(v *GetBackendInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetBackendInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetBackendJobInput(v *GetBackendJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetBackendJobInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.BackendEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackendEnvironmentName"))
	}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTokenInput(v *GetTokenInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTokenInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListBackendJobsInput(v *ListBackendJobsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListBackendJobsInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.BackendEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackendEnvironmentName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRemoveAllBackendsInput(v *RemoveAllBackendsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemoveAllBackendsInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRemoveBackendConfigInput(v *RemoveBackendConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemoveBackendConfigInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateBackendAPIInput(v *UpdateBackendAPIInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateBackendAPIInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.BackendEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackendEnvironmentName"))
	}
	if v.ResourceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateBackendAuthInput(v *UpdateBackendAuthInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateBackendAuthInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.BackendEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackendEnvironmentName"))
	}
	if v.ResourceConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceConfig"))
	} else if v.ResourceConfig != nil {
		if err := validateUpdateBackendAuthResourceConfig(v.ResourceConfig); err != nil {
			invalidParams.AddNested("ResourceConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.ResourceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateBackendConfigInput(v *UpdateBackendConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateBackendConfigInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateBackendJobInput(v *UpdateBackendJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateBackendJobInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.BackendEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackendEnvironmentName"))
	}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}