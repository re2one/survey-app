import {ChangeDetectorRef, Component, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {UserService} from '../../services/user.service';
import {SmolUser} from '../../models/smoUsers';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {QuestionsService} from '../../services/questions.service';
import {Puzzlepiece} from '../../models/puzzle';

@Component({
  selector: 'app-survey-inspect',
  templateUrl: './survey-inspect.component.html',
  styleUrls: ['./survey-inspect.component.css']
})
export class SurveyInspectComponent implements OnInit {

  surveyId: string;
  users: Array<SmolUser>;
  currentUsersAnswers: Map<any, any>;
  currentQuestions: Array<string>;
  public userForm: FormGroup;
  public questionForm: FormGroup;

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
    this.questionForm = this.formBuilder.group({
      question: ['', [Validators.required]]
    });
  }

  ngOnInit(): void {

    this.activatedRoute.paramMap.subscribe(params => {
      this.surveyId = params.get('surveyId');
      this.userService.getAll().subscribe((response) => {
        if (response.status === 200) {
          this.users = response.body;
          this.cdr.detectChanges();
        }
      });
    });
  }

  onUserFormSubmit(userEmail): void {
    this.questionsService.getAnsweredQuestions(userEmail.email).subscribe(response => {
      if (response.status === 200) {
        const map = new Map<string, Array<Puzzlepiece>>();
        this.currentQuestions = Object.keys(response.body.submissions);
        Object.keys(response.body.submissions).forEach(key => {
          const arr = new Array<Puzzlepiece>();
          response.body.submissions[key].forEach(piece => {
            const p = new Puzzlepiece();
            p.Position = piece.position;
            p.Tapped = piece.tapped;
            p.Image = piece.image;
            arr.push(p);
          });
          map.set(key, arr);
        });
        this.currentUsersAnswers = map;
      }
    });
  }

  onSubmissionSubmit(value): void {
    console.log(this.currentUsersAnswers.get(value.question));
  }
}
