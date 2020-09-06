import {Survey} from './survey';

export class Question{
  public ID: number;
  public CreatedAt: string;
  public UpdatedAt: string;
  public DeletedAt: string;
  public Survey: Survey;
  public surveyId: number;
  public title: string;
  public text: string;
  public type: string;
}

export class Questions {
  public questions: Array<Question>;
}

export class SurveyResponse {
  public question: Question;
}
