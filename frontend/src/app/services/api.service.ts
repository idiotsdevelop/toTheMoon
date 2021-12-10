import { Injectable } from '@angular/core';
import { LoginService } from './login.service';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(
    private user : LoginService
  ) { }

  // login(req: A): 
}
