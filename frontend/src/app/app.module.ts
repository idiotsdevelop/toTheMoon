import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { LoginComponent } from './login/login.component';
import { NzModule } from './nz.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

@NgModule({
  
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    NzModule,
    FormsModule,
    ReactiveFormsModule
  ],
  
  declarations: [
    AppComponent,
    LoginComponent
  ],
  
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
