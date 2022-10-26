# learn-git

## Step One: Ammend a commit message
Checkout the first step branch:
```
git checkout step-one-ammend-commit-msg
```
Run git show to see what the patch does and the bad commit message:
```
git show
```
Then run to correct the commit message:
```
git commit --amend
```

## Step Two: Rebase gainest main branch
Checkout the second step branch:
```
git checkout step-two-rebase
```
Rebase against `main` branch
```
git rebase main
```
You now should see a merge conflict that needs to be resloved.
This ca be checked by running.
```
git status
```
Once the merge has been relsoved run 
```
git rebase --continue
```
If you mess up the rebase you ran reset the to before the rebase by running
```
git rebase --abort
```

## Step Three: Perform an interactive rebase and squash commits
Checkout the third step branch:
```
git checkout step-three-squash
```
Perform an interactive rebase of the last two commits 
```
git rebase -i HEAD~2
```
Squash the head in to the pervious commit and save, now view the head commit message with
```
git log
```
or 
```
git show
```

## Step Four: Merge a feture branch
Checkout the fourth step main branch:
```
git checkout step-four-merge-main
```
Now merge the feature branch (from remote) into main branch 
```
git merge --log --no-ff origin/step-four-merge-feature
```
Fix up the conflicts and run
```
git merge --continue
```