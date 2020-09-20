import {ChangeDetectorRef, Component, OnInit} from '@angular/core';
import {Question} from '../../models/questions';
import {ActivatedRoute, Router} from '@angular/router';
import {QuestionsService} from '../../services/questions.service';
import {HttpResponse} from '@angular/common/http';
import {MuchoService} from '../../services/mucho.service';
import {SurveyResponse} from '../../models/survey';

@Component({
  selector: 'app-question-edit',
  templateUrl: './question-edit.component.html',
  styleUrls: ['./question-edit.component.css']
})
export class QuestionEditComponent implements OnInit {

  questionId: string;
  surveyId: string;
  question: Question;
  localAnswers: Map<any, any>;
  constructor(
    public router: Router,
    private questionsService: QuestionsService,
    private activatedRoute: ActivatedRoute,
    private cdr: ChangeDetectorRef,
    private answersService: MuchoService
  ) {
    this.localAnswers = new Map();
  }

  ngOnInit(): void {
    this.activatedRoute.paramMap.subscribe(params => {
      this.questionId = params.get('questionId');
      this.surveyId = params.get('surveyId');
    });
    setTimeout(() => {
      this.answersService.getAnswers(this.questionId).subscribe( (response: HttpResponse<any>) => {
        if (response.status === 200) {
          response.body.choices.forEach(answer => {
            this.localAnswers.set(answer.ID, answer);
          });
          console.log(this.localAnswers);
        }
      });
    }, 0);
  }
  permissionCheck(): boolean {
    const role = localStorage.getItem('role');
    return role === 'admin';
  }
  delete(questionId: number): void {
    console.log(`question id to be deleted: ${questionId}`);
    this.answersService.deleteAnswers(questionId).subscribe((response: HttpResponse<any>) => {
      if (response.status === 200) {
        this.localAnswers.delete(questionId);
        this.cdr.detectChanges();
      }
    });
  }
  onQuestionSubmit(questionData): void{
    this.questionsService.putQuestion(
      parseInt(this.questionId, 10),
      questionData.surveyid,
      questionData.title,
      questionData.text,
      questionData.first,
    ).subscribe((response: HttpResponse<SurveyResponse>) => {
      if (response.status === 200) {
        this.router.navigate(['/surveys/edit', questionData.surveyId]);
      }
    });
  }
  moveToEditForm(answerId: number): void {
    this.router.navigate(['/multiple/edit', answerId, this.surveyId]);
  }
  moveToAddForm(): void {
    this.router.navigate(['/multiple/add', this.questionId, this.surveyId]);
  }
}
