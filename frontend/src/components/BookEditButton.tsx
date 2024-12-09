import { HStack } from "@chakra-ui/react";
import { Button } from "@/components/ui/button";
import { AiFillEdit } from "react-icons/ai";

const BookEditButton = () => {
  return (
    <HStack>
      <Button colorPalette="teal" variant="solid">
        <AiFillEdit /> Edit
      </Button>
    </HStack>
  );
};

export default BookEditButton;
