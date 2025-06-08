import "@typespec/http";
import "@typespec/openapi";
import "../models";

using TypeSpec.Http;
using TypeSpec.OpenAPI;

namespace NoteService.Routes;

@route("/notes")
@tag("Notes")
interface Notes {
  @operationId("get-note")
  @summary("メモを取得する")
  @get
  get(@path id: int64): {
    @statusCode
    statusCode: 200;

    @body
    note: Models.Note;
  };

  @operationId("create-note")
  @summary("メモを作成する")
  @post
  create(@body body: Models.Note): Models.Note;
}
