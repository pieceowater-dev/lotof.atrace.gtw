package middleware

type contextKey string

const TokenContextKey contextKey = "token"

//func AuthDirective(ctx context.Context, next graphql.Resolver, verifyTokenFunc func(string) (bool, *pb.User, error)) (any, error) {
//	raw := graphql.GetRequestContext(ctx).Headers.Get("Authorization")
//	if raw == "" {
//		return nil, errors.New("unauthorized: token is missing")
//	}
//
//	// "Bearer <token>"
//	if len(raw) > 7 && raw[:7] == "Bearer " {
//		raw = raw[7:]
//	}
//
//	valid, user, err := verifyTokenFunc(raw)
//	if err != nil || !valid || user == nil {
//		return nil, errors.New("unauthorized: token is invalid or user not found")
//	}
//
//	log.Print("User:", user.Id, user.Username)
//
//	newCtx := context.WithValue(ctx, TokenContextKey, user)
//
//	return next(newCtx)
//}
