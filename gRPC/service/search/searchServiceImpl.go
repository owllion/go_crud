package search

import (
	"context"
	"fmt"
)

// SearchServiceImpl is the implementation of the SearchService defined in the .proto file.
type SearchServiceImpl struct{	
    UnimplementedSearchServiceServer
}

// Search is the implementation of the rpc Search method in the .proto file.
func (s *SearchServiceImpl) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
    // Here, you can implement the logic for your search, 
    // using req.Query, req.PageNumber, and req.ResultPerPage as input parameters.

    // For this example, let's assume we are just returning a dummy response:
    response := &SearchResponse{
        ReturnedMsg: "Your search for '" + req.Query + "' returned no results.",
        PageNumber : fmt.Sprintf("page number is %d", req.PageNumber),
    }

    return response, nil
}
