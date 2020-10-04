import {ChangeDetectorRef, Component, OnInit, OnChanges} from '@angular/core';
import {QuestionsService} from '../../services/questions.service';
import {Question, QuestionsResponse} from '../../models/questions';
import {HttpResponse} from '@angular/common/http';
import {ActivatedRoute, Router} from '@angular/router';
import {MuchoService} from '../../services/mucho.service';
import {PuzzleService} from '../../services/puzzle.service';
import {Mucho} from '../../models/mucho';
import {FullQuestionsService} from '../../services/full-questions.service';

@Component({
  selector: 'app-question-answer',
  templateUrl: './question-answer.component.html',
  styleUrls: ['./question-answer.component.css']
})
export class QuestionAnswerComponent implements OnInit {
  public question: Question;
  public questionId: string;
  public surveyId: string;
  constructor(
    private cdr: ChangeDetectorRef,
    private questionsService: QuestionsService,
    private fullQuestuionsService: FullQuestionsService,
    private activatedRoute: ActivatedRoute,
    private router: Router
  ) {
  }

  ngOnInit(): void {
    this.activatedRoute.paramMap.subscribe(params => {
      this.questionId = params.get('questionId');
      this.surveyId = params.get('surveyId');
      this.questionsService.getQuestion(this.questionId).subscribe((response: HttpResponse<QuestionsResponse>) => {
        if (response.status === 200) {
          this.question = response.body.question;
        }
      });
    });
  }
  onAnswerSubmit(answer): void {
    const email = localStorage.getItem('email');
    this.fullQuestuionsService.postFullQuestion(email, this.question).subscribe((response: HttpResponse<QuestionsResponse>) => {
      if (response.status === 200) {
        console.log(response.body);
        this.router.navigate(['survey', this.surveyId]);
      }
    });
  }
}
