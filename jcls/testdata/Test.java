
public class Test {
	private int privateInt0 = 0;
	int int1 = 1;
	public long publicLong16 = 16;
	public final String publicFinalString = "publicFinalString";

	private void testPrivateVoidMethod() {
		System.out.println("testPrivateVoidMethod");
	}

	public final int testPublicFinalAddMethod(int a, int b) {
		return a + b;
	}

	public static void main(String[] args) {
		if (args.length > 0) {
			System.out.println("arg0: " + args[0]);
		}
		System.out.print("Test class " + args.length + "\n");
	}
}
