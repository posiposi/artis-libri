import {
  DialogActionTrigger,
  DialogBody,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogRoot,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { useRef } from "react";
import { Button, Input, Stack } from "@chakra-ui/react";
import { Field } from "@/components/ui/field";
import { Book } from "../../types/book";

interface BookEditButtonProps {
  book: Book;
}

const BookEditButton = (book: BookEditButtonProps) => {
  const ref = useRef<HTMLInputElement>(null);
  return (
    <DialogRoot initialFocusEl={() => ref.current}>
      <DialogTrigger asChild>
        <Button variant="outline">Edit</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>書籍編集</DialogTitle>
        </DialogHeader>
        <DialogBody pb="4">
          <Stack gap="4">
            <Field label="タイトル">
              <Input ref={ref} placeholder={book.book.title} />
            </Field>
            <Field label="著者">
              <Input placeholder={book.book.author} />
            </Field>
            <Field label="ジャンル">
              <Input placeholder={book.book.genre} />
            </Field>
            <Field label="出版社">
              <Input placeholder={book.book.publisher} />
            </Field>
            <Field label="出版年">
              <Input placeholder={String(book.book.publishedAt)} />
            </Field>
            <Field label="総ページ数">
              <Input placeholder={String(book.book.totalPage)} />
            </Field>
            <Field label="金額">
              <Input placeholder={String(book.book.price)} />
            </Field>
          </Stack>
        </DialogBody>
        <DialogFooter>
          <DialogActionTrigger asChild>
            <Button variant="outline">Cancel</Button>
          </DialogActionTrigger>
          <Button>Save</Button>
        </DialogFooter>
      </DialogContent>
    </DialogRoot>
  );
};

export default BookEditButton;
