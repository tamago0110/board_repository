import React from "react";
import { READ_BOARD } from "../types";
import styles from "./MyBoard.module.css";

import { RiDeleteBinLine, RiEdit2Line } from "react-icons/ri";

import { useDispatch } from "react-redux";
import { AppDispatch } from "../../app/store";
import {
  fetchAsyncDeleteMyBoard,
  setIsOpenPutBoard,
  editMyBoardId,
  editMyBoardTitle,
  editMyBoardContent,
} from "./boardSlice";

const MyBoard: React.FC<READ_BOARD> = ({ id, created_by, title, content }) => {
  const dispatch: AppDispatch = useDispatch();

  return (
    <li className={styles.listStyle}>
      <a className={styles.iconContainer}>
        <RiDeleteBinLine
          className={styles.deleteIcon}
          onClick={() => dispatch(fetchAsyncDeleteMyBoard(id))}
        />
        <RiEdit2Line
          className={styles.editIcon}
          onClick={() => {
            dispatch(setIsOpenPutBoard());
            dispatch(editMyBoardId(id));
            dispatch(editMyBoardTitle(title));
            dispatch(editMyBoardContent(content));
          }}
        />
      </a>
      <h5 className={styles.myBoardTitle}>{title}</h5>
    </li>
  );
};

export default MyBoard;
