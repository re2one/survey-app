import {ChangeDetectorRef, Component, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {UserService} from '../../services/user.service';
import {SmolUser} from '../../models/smoUsers';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {QuestionsService} from '../../services/questions.service';
import {PuzzlePiece} from '../question-edit-puzzle/question-edit-puzzle.component';
import {ResultService} from '../../services/result.service';

@Component({
  selector: 'app-survey-inspect',
  templateUrl: './survey-inspect.component.html',
  styleUrls: ['./survey-inspect.component.css']
})
export class SurveyInspectComponent implements OnInit {

  surveyId: string;
  users: Array<SmolUser>;
  currentUsersAnswers: Map<any, any>;
  presentedPieces: Map<any, any>;
  currentQuestions: Array<string>;
  questionId: string;
  currentScore: number;
  currentUser: string;
  currentUserId: string;
  // fields: Array<number>;
  public userForm: FormGroup;
  public questionForm: FormGroup;

  constructor(
    private activatedRoute: ActivatedRoute,
    private userService: UserService,
    private formBuilder: FormBuilder,
    private questionsService: QuestionsService,
    private resultService: ResultService,
    private cdr: ChangeDetectorRef,
  ) {
    this.userForm = this.formBuilder.group({
      email: ['', [Validators.required]]
    });
    this.questionForm = this.formBuilder.group({
      question: ['', [Validators.required]]
    });
    this.presentedPieces = new Map();
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

    this.presentedPieces = new Map();
    for (let i = 0; i < 24; i++) {
      const piece = new PuzzlePiece(i.toString(10), -1);
      this.presentedPieces.set(i, piece);
    }

    this.currentUsersAnswers = new Map<string, Array<PuzzlePiece>>();
  }

  onUserFormSubmit(user): void {
    this.questionsService.getAnsweredQuestions(user.email, this.surveyId).subscribe(response => {
      if (response.status === 200) {
        this.currentUser = user.email;
        // const map = new Map<string, Array<PuzzlePiece>>();
        this.currentUsersAnswers.clear();
        this.currentQuestions = Object.keys(response.body.submissions);
        Object.keys(response.body.submissions).forEach(key => {
          const arr = new Array<PuzzlePiece>();
          response.body.submissions[key].forEach(piece => {
            const p = new PuzzlePiece(piece.position, parseInt(key, 10));
            p.tapped = piece.tapped;
            p.image = piece.image;
            p.empty = false;
            arr.push(p);
          });
          this.currentUsersAnswers.set(key, arr);
        });
        this.cdr.detectChanges();
        this.questionForm.reset();
      }
    });
  }

  onSubmissionSubmit(value): void {
    this.presentedPieces = new Map();
    for (let i = 0; i < 24; i++) {
      const piece = new PuzzlePiece(i.toString(10), value.question);
      this.presentedPieces.set(i, piece);
    }
    console.log('currentn questions');
    console.log(this.currentQuestions);
    console.log('value!');
    console.log(value);
    this.currentUsersAnswers.get(value.question).forEach((v, k) => {
      this.presentedPieces.set(parseInt(v.position, 10), v);
    });
    this.questionId = value.question;
    this.resultService.getSingleResult(this.surveyId, this.currentUser, this.questionId).subscribe(response => {
      if (response.status === 200) {
        this.currentScore = response.body;
      }
    });
    this.cdr.detectChanges();
  }
}
