import { HStack } from "@chakra-ui/react";
import { Button } from "@/components/ui/button";
import { AiFillDelete } from "react-icons/ai";

const BookDeleteButton = () => {
  return (
    <HStack>
      <Button colorPalette="teal" variant="solid">
        <AiFillDelete /> Edit
      </Button>
    </HStack>
  );
};

export default BookDeleteButton;
