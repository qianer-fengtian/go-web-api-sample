import "@typespec/http";
import "@typespec/openapi";
import "../models";

using TypeSpec.Http;
using TypeSpec.OpenAPI;

namespace NoteService.Routes;

@route("/notes")
@tag("Notes")
interface Notes {
  @operationId("list-notes")
  @summary("メモ一覧を取得する")
  @get
  list(): Models.Note[];

  @operationId("get-note")
  @summary("メモを取得する")
  @get
  get(@path id: int64): {
    @statusCode
    statusCode: 200;

    @body
    note: Models.Note;
  } | {
    @statusCode
    statusCode: 404;
  };

  @operationId("create-note")
  @summary("メモを作成する")
  @post
  create(@body body: Models.Note): Models.Note;

  @operationId("delete-note")
  @summary("メモを削除する")
  @delete
  delete(@path id: int64): {
    @statusCode
    statusCode: 204;
  } | {
    @statusCode
    statusCode: 404;
  };
}
