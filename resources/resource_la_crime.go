package resources

import (
	_ "embed"
	"encoding/json"
)

//go:embed la_crime.json
var laCrimeJSON []byte

type LACrime struct {
	Meta struct {
		View struct {
			ID                       string `json:"id"`
			Name                     string `json:"name"`
			AssetType                string `json:"assetType"`
			Attribution              string `json:"attribution"`
			AttributionLink          string `json:"attributionLink"`
			AverageRating            int    `json:"averageRating"`
			Category                 string `json:"category"`
			CreatedAt                int    `json:"createdAt"`
			Description              string `json:"description"`
			DisplayType              string `json:"displayType"`
			DownloadCount            int    `json:"downloadCount"`
			HideFromCatalog          bool   `json:"hideFromCatalog"`
			HideFromDataJSON         bool   `json:"hideFromDataJson"`
			LicenseID                string `json:"licenseId"`
			NewBackend               bool   `json:"newBackend"`
			NumberOfComments         int    `json:"numberOfComments"`
			Oid                      int    `json:"oid"`
			Provenance               string `json:"provenance"`
			PublicationAppendEnabled bool   `json:"publicationAppendEnabled"`
			PublicationDate          int    `json:"publicationDate"`
			PublicationGroup         int    `json:"publicationGroup"`
			PublicationStage         string `json:"publicationStage"`
			RowsUpdatedAt            int    `json:"rowsUpdatedAt"`
			RowsUpdatedBy            string `json:"rowsUpdatedBy"`
			TableID                  int    `json:"tableId"`
			TotalTimesRated          int    `json:"totalTimesRated"`
			ViewCount                int    `json:"viewCount"`
			ViewLastModified         int    `json:"viewLastModified"`
			ViewType                 string `json:"viewType"`
			Approvals                []struct {
				ReviewedAt            int    `json:"reviewedAt"`
				ReviewedAutomatically bool   `json:"reviewedAutomatically"`
				State                 string `json:"state"`
				SubmissionID          int    `json:"submissionId"`
				SubmissionObject      string `json:"submissionObject"`
				SubmissionOutcome     string `json:"submissionOutcome"`
				SubmittedAt           int    `json:"submittedAt"`
				TargetAudience        string `json:"targetAudience"`
				WorkflowID            int    `json:"workflowId"`
				Reviewer              struct {
					ID          string `json:"id"`
					DisplayName string `json:"displayName"`
				} `json:"reviewer"`
				SubmissionDetails struct {
					PermissionType string `json:"permissionType"`
				} `json:"submissionDetails"`
				SubmissionOutcomeApplication struct {
					EndedAt      int    `json:"endedAt"`
					FailureCount int    `json:"failureCount"`
					StartedAt    int    `json:"startedAt"`
					Status       string `json:"status"`
				} `json:"submissionOutcomeApplication"`
				Submitter struct {
					ID          string `json:"id"`
					DisplayName string `json:"displayName"`
				} `json:"submitter"`
			} `json:"approvals"`
			ClientContext struct {
				ClientContextVariables []interface{} `json:"clientContextVariables"`
				InheritedVariables     struct {
				} `json:"inheritedVariables"`
			} `json:"clientContext"`
			Columns []struct {
				ID             int    `json:"id"`
				Name           string `json:"name"`
				DataTypeName   string `json:"dataTypeName"`
				FieldName      string `json:"fieldName"`
				Position       int    `json:"position"`
				RenderTypeName string `json:"renderTypeName"`
				Format         struct {
				} `json:"format"`
				Flags          []string `json:"flags,omitempty"`
				Description    string   `json:"description,omitempty"`
				TableColumnID  int      `json:"tableColumnId,omitempty"`
				CachedContents struct {
					NonNull string `json:"non_null"`
					Largest string `json:"largest"`
					Null    string `json:"null"`
					Top     []struct {
						Item  string `json:"item"`
						Count string `json:"count"`
					} `json:"top"`
					Smallest    string `json:"smallest"`
					Count       string `json:"count"`
					Cardinality string `json:"cardinality"`
				} `json:"cachedContents,omitempty"`
			} `json:"columns"`
			Grants []struct {
				Inherited bool     `json:"inherited"`
				Type      string   `json:"type"`
				Flags     []string `json:"flags"`
			} `json:"grants"`
			License struct {
				Name      string `json:"name"`
				LogoURL   string `json:"logoUrl"`
				TermsLink string `json:"termsLink"`
			} `json:"license"`
			Metadata struct {
				Attachments []struct {
					Filename string `json:"filename"`
					AssetID  string `json:"assetId"`
					Name     string `json:"name"`
				} `json:"attachments"`
				CustomFields struct {
					CommittedUpdateFrequency struct {
						RefreshRate string `json:"Refresh rate"`
					} `json:"Committed Update Frequency"`
					LocationSpecified struct {
						DoesThisDataHaveALocationColumnYesOrNo string `json:"Does this data have a Location column? (Yes or No)"`
						WhatGeographicUnitIsTheDataCollected   string `json:"What geographic unit is the data collected?"`
					} `json:"Location Specified"`
					DataOwner struct {
						Department string `json:"Department"`
					} `json:"Data Owner"`
				} `json:"custom_fields"`
				RowLabel              string   `json:"rowLabel"`
				AvailableDisplayTypes []string `json:"availableDisplayTypes"`
			} `json:"metadata"`
			Owner struct {
				ID                    string   `json:"id"`
				DisplayName           string   `json:"displayName"`
				ProfileImageURLLarge  string   `json:"profileImageUrlLarge"`
				ProfileImageURLMedium string   `json:"profileImageUrlMedium"`
				ProfileImageURLSmall  string   `json:"profileImageUrlSmall"`
				ScreenName            string   `json:"screenName"`
				Type                  string   `json:"type"`
				Flags                 []string `json:"flags"`
			} `json:"owner"`
			Query struct {
			} `json:"query"`
			Rights      []string `json:"rights"`
			TableAuthor struct {
				ID                    string   `json:"id"`
				DisplayName           string   `json:"displayName"`
				ProfileImageURLLarge  string   `json:"profileImageUrlLarge"`
				ProfileImageURLMedium string   `json:"profileImageUrlMedium"`
				ProfileImageURLSmall  string   `json:"profileImageUrlSmall"`
				ScreenName            string   `json:"screenName"`
				Type                  string   `json:"type"`
				Flags                 []string `json:"flags"`
			} `json:"tableAuthor"`
			Tags  []string `json:"tags"`
			Flags []string `json:"flags"`
		} `json:"view"`
	} `json:"meta"`
	Data [][]interface{} `json:"data"`
}

func LoadResourceLACrime() *LACrime {
	var result LACrime
	err := json.Unmarshal(laCrimeJSON, &result)
	if err != nil {
		panic(err)
	}
	return &result
}
