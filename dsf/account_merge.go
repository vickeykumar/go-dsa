package dsf

import (
    "sort"
)

type Account struct {
    Name string
    Emails []string
}

func AccountsMerge(accounts [][]string) [][]string {
    // create a  dsf for each email and merge them if the belong to same account
    // maintain a map of emails to dsf node as well , populate the result if they belong to same dsf
    email2dset := make(map[string]*Dset)
    for _, account := range accounts {
        name := account[0]
        emails := account[1:]
        acc := Account{
            Name: name,
            Emails: emails,
            }
        var set *Dset = nil
        for _, email := range emails {
            // check if already present in map from prevous accounts 
            // and merge with current set
            // make a set for each mail and unioun them
            newset, ok  := email2dset[email]
            if !ok {
                newset = MakeSet(acc)
            }
            
            if set == nil {
                //not initialized yet
                set = newset
            } else {
                // union
                set = UnionSet(set, newset)
            }
            email2dset[email] = set // add the merge set here
        }
    }
    
    // now setroot to merged account map 
    set2mergedaccount := make(map[*Dset]Account)
    for email, set := range email2dset {
        // root set
        rootset := FindSet(set)
        acc, ok := set2mergedaccount[rootset]
        if !ok {
            // create new account for this guy name and save it
            acc = rootset.data.(Account)
            acc.Emails = []string{}
        }
        // append this mail in the email list
        acc.Emails = append(acc.Emails, email)
        set2mergedaccount[rootset] = acc // merged account
    }
    
    // populate the result from merge set
    result := make([][]string,0)
    for _, acc := range set2mergedaccount {
        tmp := []string{acc.Name}
        sort.Strings(acc.Emails)
        tmp = append(tmp, acc.Emails...)
        result = append(result, tmp)
    }
    return result
}

