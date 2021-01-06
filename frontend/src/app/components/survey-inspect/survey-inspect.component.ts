import {ChangeDetectorRef, Component, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {UserService} from '../../services/user.service';
import {SmolUser} from '../../models/smoUsers';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {QuestionsService} from '../../services/questions.service';

@Component({
  selector: 'app-survey-inspect',
  templateUrl: './survey-inspect.component.html',
  styleUrls: ['./survey-inspect.component.css']
})
export class SurveyInspectComponent implements OnInit {

  surveyId: string;
  users: Array<SmolUser>;
  public userForm: FormGroup;

  constructor(
    private activatedRoute: ActivatedRoute,
    private userService: UserService,
    private formBuilder: FormBuilder,
    private questionsService: QuestionsService,
    private cdr: ChangeDetectorRef,
  ) {
    this.userForm = this.formBuilder.group({
      email: ['', [Validators.required]]
    });
  }

  ngOnInit(): void {

    this.activatedRoute.paramMap.subscribe(params => {
      this.surveyId = params.get('surveyId');
      this.userService.getAll().subscribe((response) => {
        if (response.status === 200) {
          this.users = response.body;
          // console.log(this.users);
          this.cdr.detectChanges();
        }
      });
    });
  }

  onUserFormSubmit(userEmail): void {
    this.questionsService.getAnsweredQuestions(userEmail.email).subscribe(response => {
      if (response.status === 200) {
        console.log(response);
      }
    });
  }

}
