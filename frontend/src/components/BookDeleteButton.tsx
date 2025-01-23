import { HStack } from "@chakra-ui/react";
import { Button } from "@/components/ui/button";
import { AiFillDelete } from "react-icons/ai";
import {
  DialogActionTrigger,
  DialogBody,
  DialogCloseTrigger,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogRoot,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";

const BookDeleteButton = () => {
  return (
    <HStack>
      <DialogRoot role="alertdialog">
        <DialogTrigger asChild>
          <Button variant="outline" size="sm">
            <AiFillDelete /> Delete
          </Button>
        </DialogTrigger>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Are you sure?</DialogTitle>
          </DialogHeader>
          <DialogBody>
            <p>
              This action cannot be undone. This will permanently delete your
              account and remove your data from our systems.
            </p>
          </DialogBody>
          <DialogFooter>
            <DialogActionTrigger asChild>
              <Button variant="outline">Cancel</Button>
            </DialogActionTrigger>
            <Button colorPalette="red">Delete</Button>
          </DialogFooter>
          <DialogCloseTrigger />
        </DialogContent>
      </DialogRoot>
    </HStack>
  );
};

export default BookDeleteButton;
