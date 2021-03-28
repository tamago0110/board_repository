import { createSlice, PayloadAction, createAsyncThunk } from "@reduxjs/toolkit";
import { RootState } from "../../app/store";
import axios from "axios";
import {
  READ_BOARD,
  POST_BOARD,
  PUT_BOARD,
  POST_DISPLAY,
  READ_DISPLAY,
  POST_LEAD,
  READ_LEAD,
  BOARD_STATE,
} from "../types";

export const fetchAsyncGetBoards = createAsyncThunk(
  "board/getBoards",
  async () => {
    const res = await axios.get<READ_BOARD[]>(
      `${process.env.REACT_APP_API_URL}/board/item`,
      {
        headers: {
          Authorization: `Bearer ${localStorage.localJWT}`,
        },
      }
    );
    return res.data;
  }
);

export const fetchAsyncPostLead = createAsyncThunk(
  "board/postLead",
  async (approval: POST_LEAD) => {
    const res = await axios.post<READ_LEAD>(
      `${process.env.REACT_APP_API_URL}/user/lead`,
      approval,
      {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.localJWT}`,
        },
      }
    );
    return res.data;
  }
);

export const fetchAsyncPostDisplay = createAsyncThunk(
  "board/postApproval",
  async (display: POST_DISPLAY) => {
    const res = await axios.post<READ_DISPLAY>(
      `${process.env.REACT_APP_API_URL}/board/reject`,
      display,
      {
        headers: {
          Authorization: `Bearer ${localStorage.localJWT}`,
          "Content-Type": "application/json",
        },
      }
    );
    return res.data;
  }
);

export const fetchAsyncPostBoard = createAsyncThunk(
  "board/postBoard",
  async (board: POST_BOARD) => {
    const res = await axios.post<READ_BOARD>(
      `${process.env.REACT_APP_API_URL}/board/item`,
      board,
      {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.localJWT}`,
        },
      }
    );
    return res.data;
  }
);

export const fetchAsyncGetMyBoards = createAsyncThunk(
  "board/getMyBoards",
  async () => {
    const res = await axios.get<READ_BOARD[]>(
      `${process.env.REACT_APP_API_URL}/board/mine`,
      {
        headers: {
          Authorization: `Bearer ${localStorage.localJWT}`,
        },
      }
    );
    return res.data;
  }
);

export const fetchAsyncDeleteMyBoard = createAsyncThunk(
  "board/deleteMyBoard",
  async (id: number) => {
    const res = await axios.delete(
      `${process.env.REACT_APP_API_URL}/board/item/${id}`,
      {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.localJWT}`,
        },
      }
    );
    return id;
  }
);

export const fetchAsyncPutMyBoard = createAsyncThunk(
  "board/putMyBoard",
  async (board: PUT_BOARD) => {
    const res = await axios.put<READ_BOARD>(
      `${process.env.REACT_APP_API_URL}/board/item/${board.id}`,
      board.putBoard,
      {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.localJWT}`,
        },
      }
    );
    return res.data;
  }
);

export const initialState: BOARD_STATE = {
  isOpenPostBoard: false,
  isOpenPutBoard: false,
  postBoard: {
    title: "",
    content: "",
  },
  postLead: {
    producer: "",
  },
  postDisplay: {
    board_id: 0,
  },
  displayBoards: [{ id: 0, created_by: "", title: "", content: "" }],
  whitchBoard: 0,
  myBoards: [{ id: 0, created_by: "", title: "", content: "" }],
  updateBoard: { id: 0, putBoard: { title: "", content: "" } },
};

export const boardSlice = createSlice({
  name: "board",
  initialState,
  reducers: {
    setIsOpenPostBoard(state) {
      state.isOpenPostBoard = true;
    },
    resetIsOpenPostBoard(state) {
      state.isOpenPostBoard = false;
    },
    setIsOpenPutBoard(state) {
      state.isOpenPutBoard = true;
    },
    resetIsOpenPutBoard(state) {
      state.isOpenPutBoard = false;
      state.updateBoard = initialState.updateBoard;
    },
    editMyBoardId(state, action) {
      state.updateBoard.id = action.payload;
    },
    editMyBoardTitle(state, action) {
      state.updateBoard.putBoard.title = action.payload;
    },
    editMyBoardContent(state, action) {
      state.updateBoard.putBoard.content = action.payload;
    },
  },
  extraReducers: (builder) => {
    builder.addCase(
      fetchAsyncGetBoards.fulfilled,
      (state, action: PayloadAction<READ_BOARD[]>) => {
        return {
          ...state,
          displayBoards: action.payload,
          whitchBoard: Math.floor(Math.random() * action.payload.length),
        };
      }
    );
    builder.addCase(fetchAsyncGetBoards.rejected, () => {
      window.location.href = "/";
    });
    builder.addCase(fetchAsyncPostLead.fulfilled, (state) => {
      return {
        ...state,
        postApproval: initialState.postLead,
      };
    });
    builder.addCase(fetchAsyncPostLead.rejected, () => {
      window.location.href = "/";
    });
    builder.addCase(fetchAsyncPostDisplay.fulfilled, (state) => {
      return {
        ...state,
        postDisplay: initialState.postDisplay,
      };
    });
    builder.addCase(
      fetchAsyncPostBoard.fulfilled,
      (state, action: PayloadAction<READ_BOARD>) => {
        return {
          ...state,
          postBoard: initialState.postBoard,
          myBoards: [...state.myBoards, action.payload],
        };
      }
    );
    builder.addCase(
      fetchAsyncGetMyBoards.fulfilled,
      (state, action: PayloadAction<READ_BOARD[]>) => {
        return {
          ...state,
          myBoards: action.payload,
        };
      }
    );
    builder.addCase(
      fetchAsyncDeleteMyBoard.fulfilled,
      (state, action: PayloadAction<number>) => {
        return {
          ...state,
          myBoards: state.myBoards.filter(
            (myBoard) => myBoard.id !== action.payload
          ),
        };
      }
    );
    builder.addCase(
      fetchAsyncPutMyBoard.fulfilled,
      (state, action: PayloadAction<READ_BOARD>) => {
        return {
          ...state,
          myBoards: state.myBoards.map((myBoard) =>
            myBoard.id === action.payload.id ? action.payload : myBoard
          ),
        };
      }
    );
  },
});

export const {
  setIsOpenPostBoard,
  resetIsOpenPostBoard,
  setIsOpenPutBoard,
  resetIsOpenPutBoard,
  editMyBoardId,
  editMyBoardTitle,
  editMyBoardContent,
} = boardSlice.actions;
export const selectIsOpenPostBoard = (state: RootState) =>
  state.board.isOpenPostBoard;
export const selectIsOpenPutBoard = (state: RootState) =>
  state.board.isOpenPutBoard;
export const selectPostBoard = (state: RootState) => state.board.postBoard;
export const selectPostLead = (state: RootState) => state.board.postLead;
export const selectPostDisplay = (state: RootState) => state.board.postDisplay;
export const selectDisplayBoards = (state: RootState) =>
  state.board.displayBoards;
export const selectWhichBoard = (state: RootState) => state.board.whitchBoard;
export const selectMyBoards = (state: RootState) => state.board.myBoards;
export const selectUpdateBoard = (state: RootState) => state.board.updateBoard;

export default boardSlice.reducer;
