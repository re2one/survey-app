import { Component, OnInit } from '@angular/core';
import {AuthTestService} from '../../services/auth-test.service';
import {Observable} from 'rxjs';

@Component({
  selector: 'app-surveys',
  templateUrl: './surveys.component.html',
  styleUrls: ['./surveys.component.css']
})
export class SurveysComponent implements OnInit {
  someValue: string;

  constructor(
    private authTestService: AuthTestService
  ) { }

  ngOnInit(): void {
    this.authTestService.get().subscribe( obj => {
      this.someValue = obj.message;
    }, error => {
      console.log(error);
    });
  }

}
