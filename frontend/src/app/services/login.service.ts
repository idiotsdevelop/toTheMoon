import { Injectable } from '@angular/core';
import * as moment from 'moment';


const accessTokenKey = "E*^@&GU";
const userIdKey = "$FG@DTGW@";
const userNameKey = "@^@^$#!#";
const expiresAtKey = "$@!SFQ@#$W";

@Injectable({
  providedIn: 'root'
})
export class LoginService {
  userId: number;
  userName: string;
  expiresAt: string;
  accessToken: string = null;

  constructor() { }

  init() {
    this.accessToken = localStorage.getItem(accessTokenKey);
    this.userId = parseInt(localStorage.getItem(userIdKey));
    this.userName = localStorage.getItem(userNameKey);
    this.expiresAt = localStorage.getItem(expiresAtKey);
  }

  isAccessTokenStale() {
    const expiresAt = moment(JSON.parse(this.expiresAt) * 1000);
    return !moment().utc(false).isBefore(expiresAt);
  }

}
