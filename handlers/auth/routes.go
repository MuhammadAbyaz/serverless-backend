package auth

import "serverless/constants"

const authPrefix = constants.Prefix + "/auth"

const LoginRoute = authPrefix + "/login"

const RegisterRoute = authPrefix + "/register"
