import { Table } from "@chakra-ui/react";
// import { Book } from "../../types/book";

const BookTable = () => {
  return (
    <Table.Root size="sm">
      <Table.Header>
        <Table.Row>
          <Table.ColumnHeader>Product</Table.ColumnHeader>
          <Table.ColumnHeader>Category</Table.ColumnHeader>
          <Table.ColumnHeader>Genre</Table.ColumnHeader>
          <Table.ColumnHeader>Publisher</Table.ColumnHeader>
          <Table.ColumnHeader>Published At</Table.ColumnHeader>
          <Table.ColumnHeader>Total Page</Table.ColumnHeader>
          <Table.ColumnHeader>Progress Page</Table.ColumnHeader>
          <Table.ColumnHeader textAlign="end">Price</Table.ColumnHeader>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {items.map((item) => (
          <Table.Row key={item.id}>
            <Table.Cell>{item.title}</Table.Cell>
            <Table.Cell>{item.author}</Table.Cell>
            <Table.Cell>{item.genre}</Table.Cell>
            <Table.Cell>{item.publisher}</Table.Cell>
            <Table.Cell>{item.published_at}</Table.Cell>
            <Table.Cell>{item.total_page}</Table.Cell>
            <Table.Cell>{item.progress_page}</Table.Cell>
            <Table.Cell textAlign="end">{item.price}</Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    </Table.Root>
  );
};

const items = [
  {
    id: 1,
    title: "ガリア戦記",
    author: "Julius Caesar",
    genre: "History",
    publisher: "Shinchosha",
    published_at: "BC.50",
    total_page: 100,
    progress_page: 33,
    price: 999.99,
  },
  { id: 2, title: "Coffee Maker", author: "Home Appliances", price: 49.99 },
  { id: 3, title: "Desk Chair", author: "Furniture", price: 150.0 },
  { id: 4, title: "Smartphone", author: "Electronics", price: 799.99 },
  { id: 5, title: "Headphones", author: "Accessories", price: 199.99 },
];

export default BookTable;
