import React, { useState } from "react";
import Modal from "react-modal";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import TextField from "@material-ui/core/TextField";
import Toolbar from "@material-ui/core/Toolbar";
import Button from "@material-ui/core/Button";
import CancelIcon from "@material-ui/icons/Cancel";
import NoteAddIcon from "@material-ui/icons/NoteAdd";
import { useSelector, useDispatch } from "react-redux";
import { AppDispatch } from "../../app/store";
import { POST_BOARD } from "../types";

import {
  selectIsOpenPostBoard,
  resetIsOpenPostBoard,
  fetchAsyncPostBoard,
} from "./boardSlice";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    grow: {
      flexGrow: 1,
    },
  })
);

const customStyles = {
  content: {
    width: 400,
    height: 390,
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
  },
};

const CreateBoard: React.FC = () => {
  const classes = useStyles();
  const dispatch: AppDispatch = useDispatch();
  const isOpenPostBoard = useSelector(selectIsOpenPostBoard);
  const [newBoard, setNewBoard] = useState<POST_BOARD>({
    title: "",
    content: "",
  });
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    const name = e.target.name;
    setNewBoard({ ...newBoard, [name]: value });
  };

  return (
    <>
      <Modal
        isOpen={isOpenPostBoard}
        onRequestClose={() => dispatch(resetIsOpenPostBoard())}
        style={customStyles}
      >
        <TextField
          InputLabelProps={{
            shrink: true,
          }}
          required
          fullWidth
          label="Title"
          type="text"
          name="title"
          value={newBoard.title}
          onChange={handleInputChange}
        />
        <TextField
          InputLabelProps={{
            shrink: true,
          }}
          required
          fullWidth
          multiline
          rows={13}
          label="Content"
          type="text"
          name="content"
          value={newBoard.content}
          onChange={handleInputChange}
        />
        <Toolbar>
          <Button
            variant="contained"
            color="primary"
            onClick={() => {
              setNewBoard({ title: "", content: "" });
              dispatch(resetIsOpenPostBoard());
            }}
          >
            <CancelIcon />
            Cancel
          </Button>
          <div className={classes.grow} />
          <Button
            variant="contained"
            color="primary"
            onClick={() => {
              dispatch(fetchAsyncPostBoard(newBoard));
              setNewBoard({ title: "", content: "" });
              dispatch(resetIsOpenPostBoard());
            }}
          >
            <NoteAddIcon />
            Create
          </Button>
        </Toolbar>
      </Modal>
    </>
  );
};

export default CreateBoard;
