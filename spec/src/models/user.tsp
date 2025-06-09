namespace NoteService.Models;

model User {
  id: int64;
  name: string;
  email: string;
  createdAt?: utcDateTime;
}