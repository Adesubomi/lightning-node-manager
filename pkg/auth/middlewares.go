package auth

import (
	responsePkg "github.com/Adesubomi/lightning-node-manager/pkg/response"
	utilPkg "github.com/Adesubomi/lightning-node-manager/pkg/util"
	"github.com/gofiber/fiber/v2"
)

func (r Registry) AuthMiddleware(ctx *fiber.Ctx) error {
	bearerToken := utilPkg.GetBearerTokenFromAuthorizationHeader(ctx)

	userSession, err := r.getUserAuthSession(bearerToken)
	if err != nil || userSession == nil || userSession.User.ID == "" {
		return responsePkg.Unauthorized(ctx, "Service is only available for logged in users")
	}

	ctx.Locals("BearerToken", bearerToken)
	ctx.Locals("UserSession", userSession)

	return ctx.Next()
}

func (r Registry) GuestMiddleware(ctx *fiber.Ctx) error {
	bearerToken := utilPkg.GetBearerTokenFromAuthorizationHeader(ctx)
	userSession, err := r.getUserAuthSession(bearerToken)

	if err == nil {
		return responsePkg.Unauthorized(ctx, "Service is only available to guests")
	}

	if userSession != nil && userSession.User.ID != "" {
		return responsePkg.Unauthorized(ctx, "Service is only available to guests")
	}

	return ctx.Next()
}
