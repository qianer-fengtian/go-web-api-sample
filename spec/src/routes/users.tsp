import "@typespec/http";
import "@typespec/openapi";
import "../models";

using TypeSpec.Http;
using TypeSpec.OpenAPI;

namespace NoteService.Routes;

@route("/users")
@tag("Users")
interface Users {
  @operationId("list-users")
  @summary("ユーザー一覧を取得する")
  @get
  list(): Models.User[];

  @operationId("get-user")
  @summary("ユーザーを取得する")
  @get
  get(@path id: int64): {
    @statusCode
    statusCode: 200;

    @body
    user: Models.User;
  } | {
    @statusCode
    statusCode: 404;
  };

  @operationId("create-user")
  @summary("ユーザーを作成する")
  @post
  create(@body body: Models.User): Models.User;

  @operationId("delete-user")
  @summary("ユーザーを削除する")
  @delete
  delete(@path id: int64): {
    @statusCode
    statusCode: 204;
  } | {
    @statusCode
    statusCode: 404;
  };
}