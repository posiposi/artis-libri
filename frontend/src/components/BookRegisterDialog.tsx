import { Input, Stack } from "@chakra-ui/react";
import { Button } from "@/components/ui/button";
import { AiFillBook } from "react-icons/ai";
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
import { Field } from "@/components/ui/field";
import { useForm, SubmitHandler } from "react-hook-form";
import { Book } from "../../types/book";

const BookRegisterDialog = () => {
  const { register, handleSubmit } = useForm<Book>();
  const onSubmit: SubmitHandler<Book> = (data) => alert(data.title);
  return (
    <DialogRoot>
      <DialogTrigger asChild>
        <Button variant="outline">
          <AiFillBook /> Register
        </Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>書籍登録</DialogTitle>
        </DialogHeader>
        <form onSubmit={handleSubmit(onSubmit)}>
          <DialogBody pb="4">
            <Stack gap="4">
              <Field label="タイトル">
                <Input {...register("title")} />
              </Field>
              <Field label="著者">
                <Input {...register("author")} />
              </Field>
              <Field label="ジャンル">
                <Input {...register("genre")} />
              </Field>
              <Field label="出版社">
                <Input {...register("publisher")} />
              </Field>
              <Field label="出版年">
                <Input {...register("publishedAt")} />
              </Field>
              <Field label="総ページ数">
                <Input {...register("totalPage")} />
              </Field>
              <Field label="現状ページ">
                <Input {...register("progressPage")} placeholder="0" />
              </Field>
              <Field label="金額">
                <Input {...register("price")} />
              </Field>
            </Stack>
          </DialogBody>
          <DialogFooter>
            <DialogActionTrigger asChild>
              <Button variant="outline">Cancel</Button>
            </DialogActionTrigger>
            <Button type="submit" variant="outline" colorPalette="blue">
              Save
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </DialogRoot>
  );
};

export default BookRegisterDialog;
