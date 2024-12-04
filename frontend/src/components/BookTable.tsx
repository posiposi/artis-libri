import { useEffect, useState } from "react";
import { Table } from "@chakra-ui/react";
import { Book } from "../../types/book";
import { Box, Spinner } from "@chakra-ui/react";

const BookTable = () => {
  const [books, setBooks] = useState<Book[]>([]);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    const fetchBooks = async () => {
      try {
        const response = await fetch("http://localhost:8081/v1/books");
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
          <Table.ColumnHeader textAlign="end">金額</Table.ColumnHeader>
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
            <Table.Cell>{book.total_page}</Table.Cell>
            <Table.Cell>{book.progress_page}</Table.Cell>
            <Table.Cell textAlign="end">
              ¥{Number(book.price).toLocaleString()}
            </Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    </Table.Root>
  );
};

export default BookTable;
