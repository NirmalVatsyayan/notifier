/**
 * Created by Nirmal.Vatsyayan on 6/26/2017.
 */
import { Component, ViewChild } from '@angular/core';
import { NgForm } from '@angular/forms';
import { ServerService } from './server.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  @ViewChild('f') signupForm: NgForm;

  data = {
    header: '',
    payload: '',
    imageUrl: '',
    userQuery: '',
    notification_time: ''
  };
  submitted = false;
  constructor(private serverService: ServerService) {}

  onSubmit() {
    this.submitted = true;
    this.data.header = this.signupForm.value.data.header;
    this.data.payload = this.signupForm.value.data.payload;
    this.data.imageUrl = this.signupForm.value.data.imageUrl;
    this.data.userQuery = this.signupForm.value.data.userQuery;
    this.data.notification_time = this.signupForm.value.data.notification_time;

    this.signupForm.reset();

    this.serverService.storeNotification(this.data)
      .subscribe(
        (response) => console.log(response),
        (error) => console.log(error)
      );
  }
}
