/*

Package deferrer provides deferring defer functions to outer scopes. It is
especially useful in unit tests (leveraging Gomega) where a more complex ("It")
single test is broken down into ("By") multiple segments and the test code is
its own segment functions. If later segments need access to resources created in
previous segments, a simple defer to clean up the resource in a segment will be
too early. In such situations, simply use the Deferrer type with its Cleanup
method deferred on the same level. Then use Deferrer.Defer to defer resource
cleanup to the outer scope of where your Deferrer instance lives. The resources
will now be cleaned up only at the end of all segments or after a segment fails.

*/

package deferrer
