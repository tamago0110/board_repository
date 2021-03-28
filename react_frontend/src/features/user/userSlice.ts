import { createSlice, PayloadAction, createAsyncThunk } from "@reduxjs/toolkit";
import { RootState } from "../../app/store";
import axios from "axios";
import {
  CRED,
  JWT,
  USER_STATE,
  USER,
  LOGIN_USER,
  PROFILE,
  PUT_PROFILE,
} from "../types";

export const fetchAsyncLogin = createAsyncThunk(
  "user/login",
  async (auth: CRED) => {
    const res = await axios.post<JWT>(
      `${process.env.REACT_APP_API_URL}/login`,
      auth,
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    return res.data;
  }
);

export const fetchAsyncRegister = createAsyncThunk(
  "user/register",
  async (auth: CRED) => {
    const res = await axios.post<USER>(
      `${process.env.REACT_APP_API_URL}/signup`,
      auth,
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    return res.data;
  }
);

export const fetchAsyncGetMyProf = createAsyncThunk(
  "user/loginuser",
  async () => {
    const res = await axios.get<LOGIN_USER>(
      `${process.env.REACT_APP_API_URL}/user/me`,
      {
        headers: {
          Authorization: `Bearer ${localStorage.localJWT}`,
        },
      }
    );
    return res.data;
  }
);

export const fetchAsyncCreateProf = createAsyncThunk(
  "user/createProfile",
  async () => {
    const res = await axios.post<PROFILE>(
      `${process.env.REACT_APP_API_URL}/user/profile`,
      { image: null },
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

export const fetchAsyncGetProfs = createAsyncThunk(
  "user/getProfiles",
  async () => {
    const res = await axios.get<PROFILE[]>(
      `${process.env.REACT_APP_API_URL}/user/profile`,
      {
        headers: {
          Authorization: `Bearer ${localStorage.localJWT}`,
        },
      }
    );
    return res.data;
  }
);

export const fetchAsyncGetSpecificProf = createAsyncThunk(
  "user/getProfile",
  async (id: string) => {
    const res = await axios.get<PROFILE>(
      `${process.env.REACT_APP_API_URL}/user/profile/${id}`,
      {
        headers: {
          Authorization: `Bearer ${localStorage.localJWT}`,
        },
      }
    );
    return res.data;
  }
);

export const fetchAsyncUpdateProfs = createAsyncThunk(
  "user/updateProfile",
  async (profile: PUT_PROFILE) => {
    const uploadData = new FormData();
    uploadData.append("name", profile.putProfile.name);
    profile.putProfile.image &&
      uploadData.append(
        "image",
        profile.putProfile.image,
        profile.putProfile.image.name
      );
    const res = await axios.put<PROFILE>(
      `${process.env.REACT_APP_API_URL}/user/me`,
      uploadData,
      {
        headers: {
          "Content-Type": "multipart/form-data",
          Authorization: `Bearer ${localStorage.localJWT}`,
        },
      }
    );
    return res.data;
  }
);

const initialState: USER_STATE = {
  isOpenPutProfile: false,
  isLoginView: true,
  loginUser: {
    id: 0,
    user_id: "",
    name: "",
    image: "",
  },
  profiles: [{ id: 0, user_id: "", name: "", image: "" }],
  updateProfile: { id: 0, putProfile: { name: "", image: null } },
};

export const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    toggleMode(state) {
      state.isLoginView = !state.isLoginView;
    },
    setIsOpenPutProfile(state) {
      state.isOpenPutProfile = true;
    },
    resetIsOpenPutProfile(state) {
      state.isOpenPutProfile = false;
      state.updateProfile.putProfile.name =
        initialState.updateProfile.putProfile.name;
      state.updateProfile.putProfile.image =
        initialState.updateProfile.putProfile.image;
    },
    editProfileName(state, action) {
      state.updateProfile.putProfile.name = action.payload;
    },
    editProfileImage(state, action) {
      state.updateProfile.putProfile.image = action.payload;
    },
  },
  extraReducers: (builder) => {
    builder.addCase(
      fetchAsyncLogin.fulfilled,
      (state, action: PayloadAction<JWT>) => {
        localStorage.setItem("localJWT", action.payload.token);
        action.payload.token && (window.location.href = "/boards");
      }
    );
    builder.addCase(
      fetchAsyncGetMyProf.fulfilled,
      (state, action: PayloadAction<LOGIN_USER>) => {
        state.loginUser = action.payload;
        state.updateProfile.id = action.payload.id;
      }
    );
    builder.addCase(
      fetchAsyncGetProfs.fulfilled,
      (state, action: PayloadAction<PROFILE[]>) => {
        return {
          ...state,
          profiles: action.payload,
        };
      }
    );
    builder.addCase(
      fetchAsyncUpdateProfs.fulfilled,
      (state, action: PayloadAction<PROFILE>) => {
        return {
          ...state,
          loginUser: action.payload,
        };
      }
    );
  },
});

export const {
  toggleMode,
  setIsOpenPutProfile,
  resetIsOpenPutProfile,
  editProfileName,
  editProfileImage,
} = userSlice.actions;

export const selectIsOpenPutProfile = (state: RootState) =>
  state.user.isOpenPutProfile;
export const selectIsLoginView = (state: RootState) => state.user.isLoginView;
export const selectLoginUser = (state: RootState) => state.user.loginUser;
export const selectProfiles = (state: RootState) => state.user.profiles;
export const selectUpdateProfile = (state: RootState) =>
  state.user.updateProfile;

export default userSlice.reducer;
