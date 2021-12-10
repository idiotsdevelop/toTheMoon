import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ApiService } from '../services/api.service';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  validateForm!: FormGroup;

  submitForm(): void {
    
  }

  constructor(
    private fb : FormBuilder,
    private apiService : ApiService,
  ) {}

  ngOnInit(): void {
    this.validateForm = this.fb.group({
      userId : [null, [Validators.required]],
      password : [null,[Validators.required]],
    })
  }

}
