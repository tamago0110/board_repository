import { FILE } from "dns";

/*userSlice*/
export interface LOGIN_USER {
  id: number;
  user_id: string;
  name: string;
  image: string;
}

export interface FILE {
  readonly lastModified: number;
  readonly name: string;
}

export interface CRED {
  email: string;
  password: string;
}

export interface JWT {
  token: string;
}

export interface USER {
  uuid: string;
  email: string;
}

export interface PROFILE {
  id: number;
  user_id: string;
  name: string;
  image: string;
}

export interface POST_PROFILE {
  name: string;
  image: File | null;
}

export interface PUT_PROFILE {
  id: number;
  putProfile: POST_PROFILE;
}

export interface USER_STATE {
  isOpenPutProfile: boolean;
  isLoginView: boolean;
  loginUser: LOGIN_USER;
  profiles: PROFILE[];
  updateProfile: PUT_PROFILE;
}

/*boardSlice*/
export interface READ_BOARD {
  id: number;
  created_by: string;
  title: string;
  content: string;
}

export interface POST_BOARD {
  title: string;
  content: string;
}

export interface PUT_BOARD {
  id: number;
  putBoard: POST_BOARD;
}

export interface READ_DISPLAY {
  id: number;
  board_id: number;
  rejected_by: string;
}

export interface POST_DISPLAY {
  board_id: number;
}

export interface POST_LEAD {
  producer: string;
}

export interface READ_LEAD {
  id: number;
  consumer: string;
  producer: string;
}

export interface BOARD_STATE {
  isOpenPostBoard: boolean;
  isOpenPutBoard: boolean;
  postBoard: POST_BOARD;
  postLead: POST_LEAD;
  postDisplay: POST_DISPLAY;
  displayBoards: READ_BOARD[];
  whitchBoard: number;
  myBoards: READ_BOARD[];
  updateBoard: PUT_BOARD;
}
