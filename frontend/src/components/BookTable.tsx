import { useEffect, useState } from "react";
import { Table } from "@chakra-ui/react";
import { Book } from "../../types/book";
import { Box, Spinner } from "@chakra-ui/react";
import BookEditButton from "./BookEditButton";
import BookDeleteButton from "./BookDeleteButton";

const BookTable = () => {
  const [books, setBooks] = useState<Book[]>([]);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    const fetchBooks = async () => {
      try {
        const baseURL = import.meta.env.VITE_API_BASE_URL;
        const response = await fetch(`${baseURL}/v1/books`);
        const data = await response.json();
        setBooks(data);
      } catch (error) {
        console.error("Error fetching books:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchBooks();
  }, []);

  if (loading) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        height="100vh"
      >
        <Spinner size="xl" />
      </Box>
    );
  }

  return (
    <Table.Root size="sm">
      <Table.Header>
        <Table.Row>
          <Table.ColumnHeader>タイトル</Table.ColumnHeader>
          <Table.ColumnHeader>著者</Table.ColumnHeader>
          <Table.ColumnHeader>ジャンル</Table.ColumnHeader>
          <Table.ColumnHeader>出版社</Table.ColumnHeader>
          <Table.ColumnHeader>出版年</Table.ColumnHeader>
          <Table.ColumnHeader>総ページ数</Table.ColumnHeader>
          <Table.ColumnHeader>現状ページ</Table.ColumnHeader>
          <Table.ColumnHeader>進捗率</Table.ColumnHeader>
          <Table.ColumnHeader>金額</Table.ColumnHeader>
          <Table.ColumnHeader></Table.ColumnHeader>
          <Table.ColumnHeader textAlign="end"></Table.ColumnHeader>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {books.map((book) => (
          <Table.Row key={book.id}>
            <Table.Cell>{book.title}</Table.Cell>
            <Table.Cell>{book.author}</Table.Cell>
            <Table.Cell>{book.genre}</Table.Cell>
            <Table.Cell>{book.publisher}</Table.Cell>
            <Table.Cell>{book.published_at}年</Table.Cell>
            <Table.Cell>{book.total_page}p</Table.Cell>
            <Table.Cell>{book.progress_page}p</Table.Cell>
            <Table.Cell>{book.progress_percentage}%</Table.Cell>
            <Table.Cell>¥{Number(book.price).toLocaleString()}</Table.Cell>
            <Table.Cell>
              <BookEditButton></BookEditButton>
            </Table.Cell>
            <Table.Cell textAlign="end">
              <BookDeleteButton></BookDeleteButton>
            </Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    </Table.Root>
  );
};

export default BookTable;
